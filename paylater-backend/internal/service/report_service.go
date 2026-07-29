package service

import (
	"context"
	"database/sql"
	
	db "paylater-backend/internal/db"
)

type ReportService struct {
	queries *db.Queries
}

func NewReportService(queries *db.Queries) *ReportService {
	return &ReportService{
		queries: queries,
	}
}

// Users who reached credit limit
func (s *ReportService) GetUsersAtCreditLimit(ctx context.Context) ([]db.Customer, error) {
	return s.queries.GetUsersAtCreditLimit(ctx)
}

// Customers having pending due
func (s *ReportService) GetCustomersWithDue(ctx context.Context) ([]db.Customer, error) {
	return s.queries.GetCustomersWithDue(ctx)
}

// Customer due by name
func (s *ReportService) GetCustomerDueByName(ctx context.Context, name string) ([]sql.NullString, error) {
	return s.queries.GetCustomerDueByName(ctx, name)
}

// Merchant fee collected
func (s *ReportService) GetAllMerchantsFeeCollected(ctx context.Context) ([]db.GetAllMerchantsFeeCollectedRow, error) {
	return s.queries.GetAllMerchantsFeeCollected(ctx)
}