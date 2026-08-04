package ledger

import (
	"context"

	db "paylater-backend/internal/db"
)

// MerchantCommissionSQLC adapts SQLC merchant queries to MerchantCommissionPort.
type MerchantCommissionSQLC struct {
	q *db.Queries
}

func NewMerchantCommissionSQLC(q *db.Queries) MerchantCommissionPort {
	return &MerchantCommissionSQLC{q: q}
}

func (a *MerchantCommissionSQLC) GetCommission(ctx context.Context, merchantID int32) (MerchantCommission, error) {
	merchant, err := a.q.GetMerchantByID(ctx, merchantID)
	if err != nil {
		return MerchantCommission{}, err
	}
	return MerchantCommission{
		ID:                   merchant.ID,
		CommissionPercentage: merchant.CommissionPercentage,
	}, nil
}
