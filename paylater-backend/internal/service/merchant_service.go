package service

import (
	"context"

	db "paylater-backend/internal/db"
)

type MerchantService struct {
	queries *db.Queries
}

func NewMerchantService(queries *db.Queries) *MerchantService {
	return &MerchantService{
		queries: queries,
	}
}

// Create Merchant
func (s *MerchantService) CreateMerchant(ctx context.Context, arg db.CreateMerchantParams) error {

	_, err := s.queries.CreateMerchant(ctx, arg)
	if err != nil {
		return err
	}

	return nil
}

// Get Merchant By ID
func (s *MerchantService) GetMerchantByID(ctx context.Context, id int32) (db.Merchant, error) {

	return s.queries.GetMerchantByID(ctx, id)
}

// Get All Merchants
func (s *MerchantService) GetAllMerchants(ctx context.Context) ([]db.Merchant, error) {

	return s.queries.GetAllMerchants(ctx)
}

// Update Merchant Commission
func (s *MerchantService) UpdateMerchantCommission(ctx context.Context, arg db.UpdateMerchantCommissionParams) error {

	return s.queries.UpdateMerchantCommission(ctx, arg)
}

//update merchant name and email
func (s *MerchantService) UpdateMerchant(ctx context.Context, arg db.UpdateMerchantParams) error {
	return s.queries.UpdateMerchant(ctx,arg)
}

func (s *MerchantService) DeleteMerchant(ctx context.Context, id int32) error {

	err := s.queries.DeleteMerchantById(ctx, id)
	if err != nil {
		return err
	}

	return nil
}