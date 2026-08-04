package ledger

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	db "paylater-backend/internal/db"
)

type Service struct {
	queries *db.Queries
}

func NewService(queries *db.Queries) *Service {
	return &Service{queries: queries}
}

func (s *Service) Purchase(ctx context.Context, customerID int32, req PurchaseRequest) error {
	params := db.CreateTransactionParams{
		CustomerID: customerID,
		MerchantID: sql.NullInt32{
			Int32: req.MerchantID,
			Valid: true,
		},
		TransactionType: db.TransactionsTransactionTypePURCHASE,
		Amount:          req.Amount,
	}
	return s.CreateTransaction(ctx, params)
}

func (s *Service) Payback(ctx context.Context, customerID int32, req PaybackRequest) error {
	params := db.CreateTransactionParams{
		CustomerID:      customerID,
		TransactionType: db.TransactionsTransactionTypePAYBACK,
		Amount:          req.Amount,
	}
	return s.CreateTransaction(ctx, params)
}

func (s *Service) CreateTransaction(ctx context.Context, arg db.CreateTransactionParams) error {
	customer, err := s.queries.GetCustomerByID(ctx, arg.CustomerID)
	if err != nil {
		return errors.New("customer not found")
	}

	if customer.Status.Valid &&
		customer.Status.CustomersStatus == db.CustomersStatusBLOCKED {
		return errors.New("customer is blocked")
	}

	amount, err := strconv.ParseFloat(arg.Amount, 64)
	if err != nil {
		return errors.New("invalid amount")
	}
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	creditLimit, err := strconv.ParseFloat(customer.CreditLimit, 64)
	if err != nil {
		return err
	}

	currentDue := 0.0
	if customer.TotalDue.Valid {
		currentDue, err = strconv.ParseFloat(customer.TotalDue.String, 64)
		if err != nil {
			return err
		}
	}

	switch arg.TransactionType {
	case db.TransactionsTransactionTypePURCHASE:
		if time.Now().After(customer.PaymentDueDate) {
			err := s.queries.UpdateCustomerStatus(
				ctx,
				db.UpdateCustomerStatusParams{
					ID: customer.ID,
					Status: db.NullCustomersStatus{
						CustomersStatus: db.CustomersStatusBLOCKED,
						Valid:           true,
					},
				},
			)
			if err != nil {
				return err
			}
			return errors.New("payment due date crossed, customer has been blocked")
		}

		if !arg.MerchantID.Valid {
			return errors.New("merchant id required")
		}

		merchant, err := s.queries.GetMerchantByID(ctx, arg.MerchantID.Int32)
		if err != nil {
			return errors.New("merchant not found")
		}

		if currentDue+amount > creditLimit {
			return errors.New("credit limit exceeded")
		}

		commissionPercentage := 0.0
		if merchant.CommissionPercentage.Valid {
			commissionPercentage, err = strconv.ParseFloat(
				merchant.CommissionPercentage.String,
				64,
			)
			if err != nil {
				return err
			}
		}

		commissionAmount := (amount * commissionPercentage) / 100

		arg.CommissionPercentage = merchant.CommissionPercentage
		arg.CommissionAmount = sql.NullString{
			String: fmt.Sprintf("%.2f", commissionAmount),
			Valid:  true,
		}

		_, err = s.queries.CreateTransaction(ctx, arg)
		if err != nil {
			return err
		}

		newDue := currentDue + amount
		return s.queries.UpdateCustomerDue(ctx, db.UpdateCustomerDueParams{
			ID: customer.ID,
			TotalDue: sql.NullString{
				String: fmt.Sprintf("%.2f", newDue),
				Valid:  true,
			},
		})

	case db.TransactionsTransactionTypePAYBACK:
		if currentDue == 0 {
			return errors.New("customer has no outstanding due")
		}
		if amount > currentDue {
			return errors.New("payback amount exceeds current due")
		}

		arg.MerchantID = sql.NullInt32{}
		arg.CommissionPercentage = sql.NullString{}
		arg.CommissionAmount = sql.NullString{}

		_, err = s.queries.CreateTransaction(ctx, arg)
		if err != nil {
			return err
		}

		newDue := currentDue - amount
		return s.queries.UpdateCustomerDue(ctx, db.UpdateCustomerDueParams{
			ID: customer.ID,
			TotalDue: sql.NullString{
				String: fmt.Sprintf("%.2f", newDue),
				Valid:  true,
			},
		})

	default:
		return errors.New("invalid transaction type")
	}
}

func (s *Service) GetAllTransactions(ctx context.Context) ([]db.Transaction, error) {
	return s.queries.GetAllTransactions(ctx)
}

func (s *Service) GetTransactionsByCustomerID(ctx context.Context, customerID int32) ([]db.Transaction, error) {
	return s.queries.GetTransactionsByCustomerID(ctx, customerID)
}

func (s *Service) GetTransactionsByMerchantID(ctx context.Context, merchantID sql.NullInt32) ([]db.Transaction, error) {
	return s.queries.GetTransactionsByMerchantID(ctx, merchantID)
}
