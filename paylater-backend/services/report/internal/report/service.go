package report

import (
	"context"

	db "paylater/services/report/internal/db"
)

type Service struct {
	queries *db.Queries
}

func NewService(queries *db.Queries) *Service {
	return &Service{queries: queries}
}

func (s *Service) GetUsersAtCreditLimit(ctx context.Context) ([]db.Customer, error) {
	return s.queries.GetUsersAtCreditLimit(ctx)
}

func (s *Service) GetCustomersWithDue(ctx context.Context) ([]db.Customer, error) {
	return s.queries.GetCustomersWithDue(ctx)
}

func (s *Service) GetCustomerDueByName(ctx context.Context, name string) (db.Customer, error) {
	return s.queries.GetCustomerDueByName(ctx, name)
}

func (s *Service) GetAllMerchantsFeeCollected(ctx context.Context) ([]db.GetAllMerchantsFeeCollectedRow, error) {
	return s.queries.GetAllMerchantsFeeCollected(ctx)
}
