package ledger

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	db "paylater/services/ledger/internal/db"
)

type Service struct {
	db        *sql.DB
	queries   *db.Queries
	customers CustomerCreditPort
	merchants MerchantCommissionPort
}

func NewService(
	database *sql.DB,
	queries *db.Queries,
	customers CustomerCreditPort,
	merchants MerchantCommissionPort,
) *Service {
	return &Service{
		db:        database,
		queries:   queries,
		customers: customers,
		merchants: merchants,
	}
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
	// Resolve remote credit/commission before opening the local DB transaction
	// so we do not hold MySQL locks across HTTP calls.
	customer, err := s.customers.GetForUpdate(ctx, arg.CustomerID)
	if err != nil {
		if err.Error() == "customer not found" {
			return errors.New("customer not found")
		}
		return err
	}

	if customer.StatusValid && customer.Status == "BLOCKED" {
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
		if dateOnly(time.Now()).After(dateOnly(customer.PaymentDueDate)) {
			if err := s.customers.Block(ctx, customer.ID); err != nil {
				return err
			}
			return errors.New("payment due date crossed, customer has been blocked")
		}

		if !arg.MerchantID.Valid {
			return errors.New("merchant id required")
		}

		merchant, err := s.merchants.GetCommission(ctx, arg.MerchantID.Int32)
		if err != nil {
			if err.Error() == "merchant not found" {
				return errors.New("merchant not found")
			}
			return err
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

		// Update customer due over HTTP before the local ledger write.
		// Holding a ledger TX open across HTTP deadlocks when transactions
		// and customers share MySQL (FK / row locks on customers).
		newDue := currentDue + amount
		if err := s.customers.UpdateDue(ctx, customer.ID, fmt.Sprintf("%.2f", newDue)); err != nil {
			return err
		}

		tx, err := s.db.BeginTx(ctx, nil)
		if err != nil {
			_ = s.customers.UpdateDue(ctx, customer.ID, fmt.Sprintf("%.2f", currentDue))
			return err
		}
		defer tx.Rollback()

		if _, err = s.queries.WithTx(tx).CreateTransaction(ctx, arg); err != nil {
			_ = s.customers.UpdateDue(ctx, customer.ID, fmt.Sprintf("%.2f", currentDue))
			return err
		}

		if err := tx.Commit(); err != nil {
			_ = s.customers.UpdateDue(ctx, customer.ID, fmt.Sprintf("%.2f", currentDue))
			return err
		}
		return nil

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

		newDue := currentDue - amount
		if err := s.customers.UpdateDue(ctx, customer.ID, fmt.Sprintf("%.2f", newDue)); err != nil {
			return err
		}

		tx, err := s.db.BeginTx(ctx, nil)
		if err != nil {
			_ = s.customers.UpdateDue(ctx, customer.ID, fmt.Sprintf("%.2f", currentDue))
			return err
		}
		defer tx.Rollback()

		if _, err = s.queries.WithTx(tx).CreateTransaction(ctx, arg); err != nil {
			_ = s.customers.UpdateDue(ctx, customer.ID, fmt.Sprintf("%.2f", currentDue))
			return err
		}

		if err := tx.Commit(); err != nil {
			_ = s.customers.UpdateDue(ctx, customer.ID, fmt.Sprintf("%.2f", currentDue))
			return err
		}
		return nil

	default:
		return errors.New("invalid transaction type")
	}
}

func dateOnly(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
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

func (s *Service) GetMerchantFeeTotals(ctx context.Context) ([]MerchantFeeTotal, error) {
	rows, err := s.queries.GetMerchantFeeTotals(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]MerchantFeeTotal, 0, len(rows))
	for _, row := range rows {
		if !row.MerchantID.Valid {
			continue
		}
		out = append(out, MerchantFeeTotal{
			MerchantID:        row.MerchantID.Int32,
			TotalFeeCollected: feeCollectedToString(row.TotalFeeCollected),
		})
	}
	return out, nil
}

func feeCollectedToString(v interface{}) string {
	if v == nil {
		return "0"
	}
	switch t := v.(type) {
	case []byte:
		if len(t) == 0 {
			return "0"
		}
		return string(t)
	case string:
		if t == "" {
			return "0"
		}
		return t
	default:
		return fmt.Sprintf("%v", t)
	}
}
