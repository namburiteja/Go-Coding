package service

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"fmt"
	db "paylater-backend/internal/db"
)

type TransactionService struct {
	queries *db.Queries
}

func NewTransactionService(queries *db.Queries) *TransactionService {
	return &TransactionService{
		queries: queries,
	}
}

func (s *TransactionService) CreateTransaction(ctx context.Context, arg db.CreateTransactionParams) error {

	// Get Customer
	customer, err := s.queries.GetCustomerByID(ctx, arg.CustomerID)
	if err != nil {
		return errors.New("customer not found")
	}

	// Convert Amount
	amount, err := strconv.ParseFloat(arg.Amount, 64)
	if err != nil {
		return errors.New("invalid amount")
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

		// Merchant Required
		if !arg.MerchantID.Valid {
			return errors.New("merchant id required")
		}

		// Merchant Exists?
		merchant, err := s.queries.GetMerchantByID(ctx, arg.MerchantID.Int32)
		if err != nil {
			return errors.New("merchant not found")
		}

		// Credit Limit Check
		if currentDue+amount > creditLimit {
			return errors.New("credit limit exceeded")
		}

		// Commission %
		commissionPercentage, err := strconv.ParseFloat(merchant.CommissionPercentage, 64)
		if err != nil {
			return err
		}

		commissionAmount := (amount * commissionPercentage) / 100

		arg.CommissionPercentage = sql.NullString{
			String: merchant.CommissionPercentage,
			Valid:  true,
		}

		arg.CommissionAmount = sql.NullString{
			String: fmt.Sprintf("%.2f", commissionAmount),
			Valid:  true,
		}

		// Save Transaction
		_, err = s.queries.CreateTransaction(ctx, arg)
		if err != nil {
			return err
		}

		// Update Due
		newDue := currentDue + amount

		return s.queries.UpdateCustomerDue(ctx, db.UpdateCustomerDueParams{
			ID: customer.ID,
			TotalDue: sql.NullString{
				String: fmt.Sprintf("%.2f", newDue),
				Valid:  true,
			},
		})

	case db.TransactionsTransactionTypePAYBACK:

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