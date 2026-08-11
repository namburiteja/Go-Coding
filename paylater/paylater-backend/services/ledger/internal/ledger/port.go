package ledger

import (
	"context"
	"database/sql"
	"time"
)

// CreditAccount is the customer credit state Ledger needs for purchase/payback.
type CreditAccount struct {
	ID             int32
	CreditLimit    string
	TotalDue       sql.NullString
	PaymentDueDate time.Time
	Status         string
	StatusValid    bool
}

// MerchantCommission is the merchant data Ledger needs for purchase fees.
type MerchantCommission struct {
	ID                   int32
	CommissionPercentage sql.NullString
}

// CustomerCreditPort is the cross-domain seam Ledger uses for credit operations.
type CustomerCreditPort interface {
	GetForUpdate(ctx context.Context, customerID int32) (CreditAccount, error)
	UpdateDue(ctx context.Context, customerID int32, totalDue string) error
	Block(ctx context.Context, customerID int32) error
}

// MerchantCommissionPort is the cross-domain seam Ledger uses for commission reads.
type MerchantCommissionPort interface {
	GetCommission(ctx context.Context, merchantID int32) (MerchantCommission, error)
}
