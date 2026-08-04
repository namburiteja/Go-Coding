package ledger

import (
	"context"
	"database/sql"

	db "paylater-backend/internal/db"
)

// CustomerCreditSQLC adapts SQLC customer queries to CustomerCreditPort.
type CustomerCreditSQLC struct {
	q *db.Queries
}

func NewCustomerCreditSQLC(q *db.Queries) CustomerCreditPort {
	return &CustomerCreditSQLC{q: q}
}

func (a *CustomerCreditSQLC) GetForUpdate(ctx context.Context, customerID int32) (CreditAccount, error) {
	customer, err := a.q.GetCustomerByIDForUpdate(ctx, customerID)
	if err != nil {
		return CreditAccount{}, err
	}
	return CreditAccount{
		ID:             customer.ID,
		CreditLimit:    customer.CreditLimit,
		TotalDue:       customer.TotalDue,
		PaymentDueDate: customer.PaymentDueDate,
		Status:         customer.Status,
	}, nil
}

func (a *CustomerCreditSQLC) UpdateDue(ctx context.Context, customerID int32, totalDue string) error {
	return a.q.UpdateCustomerDue(ctx, db.UpdateCustomerDueParams{
		ID: customerID,
		TotalDue: sql.NullString{
			String: totalDue,
			Valid:  true,
		},
	})
}

func (a *CustomerCreditSQLC) Block(ctx context.Context, customerID int32) error {
	return a.q.UpdateCustomerStatus(ctx, db.UpdateCustomerStatusParams{
		ID: customerID,
		Status: db.NullCustomersStatus{
			CustomersStatus: db.CustomersStatusBLOCKED,
			Valid:           true,
		},
	})
}
