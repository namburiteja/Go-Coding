package service

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"
	"fmt"
	db "paylater-backend/internal/db"
	"paylater-backend/internal/dto"
)

type TransactionService struct {
	queries *db.Queries
}

func NewTransactionService(queries *db.Queries) *TransactionService {
	return &TransactionService{
		queries: queries,
	}
}

func (s *TransactionService) Purchase(
	ctx context.Context,
	customerID int32,
	req dto.PurchaseRequest,
) error {

	params := db.CreateTransactionParams{
		CustomerID: customerID,

		MerchantID: sql.NullInt32{
			Int32: req.MerchantID,
			Valid: true,
		},

		TransactionType: db.TransactionsTransactionTypePURCHASE,

		Amount: req.Amount,
	}

	return s.CreateTransaction(ctx, params)
}
func (s *TransactionService) Payback(
	ctx context.Context,
	customerID int32,
	req dto.PaybackRequest,
) error {

	params := db.CreateTransactionParams{
		CustomerID: customerID,

		TransactionType: db.TransactionsTransactionTypePAYBACK,

		Amount: req.Amount,
	}

	return s.CreateTransaction(ctx, params)
}

func (s *TransactionService) CreateTransaction(ctx context.Context, arg db.CreateTransactionParams) error {

	// Get Customer
	customer, err := s.queries.GetCustomerByID(ctx, arg.CustomerID)
	if err != nil {
		return errors.New("customer not found")
	}

	// Customer Status Check
	if customer.Status.Valid &&
		customer.Status.CustomersStatus == db.CustomersStatusBLOCKED {

		return errors.New("customer is blocked")
	}

	// Convert Amount
	amount, err := strconv.ParseFloat(arg.Amount, 64)
	if err != nil {
		return errors.New("invalid amount")
	}

	// Amount Validation
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	// Convert Credit Limit
	creditLimit, err := strconv.ParseFloat(customer.CreditLimit, 64)
	if err != nil {
		return err
	}

	// Convert Current Due
	currentDue := 0.0

	if customer.TotalDue.Valid {
		currentDue, err = strconv.ParseFloat(customer.TotalDue.String, 64)
		if err != nil {
			return err
		}
	}

	switch arg.TransactionType {

	case db.TransactionsTransactionTypePURCHASE:

		// Due Date Check
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

		// Merchant Required
		if !arg.MerchantID.Valid {
			return errors.New("merchant id required")
		}

		// Merchant Exists
		merchant, err := s.queries.GetMerchantByID(ctx, arg.MerchantID.Int32)
		if err != nil {
			return errors.New("merchant not found")
		}

		// Credit Limit Check
		if currentDue+amount > creditLimit {
			return errors.New("credit limit exceeded")
		}

		// Commission Percentage
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

		// Save Transaction
		_, err = s.queries.CreateTransaction(ctx, arg)
		if err != nil {
			return err
		}

		// Update Customer Due
		newDue := currentDue + amount

		return s.queries.UpdateCustomerDue(ctx, db.UpdateCustomerDueParams{
			ID: customer.ID,
			TotalDue: sql.NullString{
				String: fmt.Sprintf("%.2f", newDue),
				Valid:  true,
			},
		})

	case db.TransactionsTransactionTypePAYBACK:

		// No Due
		if currentDue == 0 {
			return errors.New("customer has no outstanding due")
		}

		// Cannot Pay More Than Due
		if amount > currentDue {
			return errors.New("payback amount exceeds current due")
		}

		arg.MerchantID = sql.NullInt32{}

		arg.CommissionPercentage = sql.NullString{}

		arg.CommissionAmount = sql.NullString{}

		// Save Transaction
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

func (s *TransactionService) GetAllTransactions(ctx context.Context) ([]db.Transaction, error) {
	return s.queries.GetAllTransactions(ctx)
}

func (s *TransactionService) GetTransactionsByCustomerID(ctx context.Context, customerID int32) ([]db.Transaction, error) {
	return s.queries.GetTransactionsByCustomerID(ctx, customerID)
}

func (s *TransactionService) GetTransactionsByMerchantID(ctx context.Context, merchantID sql.NullInt32) ([]db.Transaction, error) {
	return s.queries.GetTransactionsByMerchantID(ctx, merchantID)
}