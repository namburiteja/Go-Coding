package ledger

import (
	"context"
	"database/sql"
	"time"

	db "paylater-backend/internal/db"
)

// CreditAccount is the customer credit state Ledger needs for purchase/payback.
type CreditAccount struct {
	ID             int32
	CreditLimit    string
	TotalDue       sql.NullString
	PaymentDueDate time.Time
	Status         db.NullCustomersStatus
}

// MerchantCommission is the merchant data Ledger needs for purchase fees.
type MerchantCommission struct {
	ID                   int32
	CommissionPercentage sql.NullString
}

// CustomerCreditPort is the cross-domain seam Ledger uses for credit operations.
// In-process SQLC today; HTTP in a later phase.
type CustomerCreditPort interface {
	GetForUpdate(ctx context.Context, customerID int32) (CreditAccount, error)
	UpdateDue(ctx context.Context, customerID int32, totalDue string) error
	Block(ctx context.Context, customerID int32) error
}

// MerchantCommissionPort is the cross-domain seam Ledger uses for commission reads.
type MerchantCommissionPort interface {
	GetCommission(ctx context.Context, merchantID int32) (MerchantCommission, error)
}

// CustomerCreditFactory builds a CustomerCreditPort bound to a Queries (tx or not).
type CustomerCreditFactory func(q *db.Queries) CustomerCreditPort

// MerchantCommissionFactory builds a MerchantCommissionPort bound to a Queries (tx or not).
type MerchantCommissionFactory func(q *db.Queries) MerchantCommissionPort
