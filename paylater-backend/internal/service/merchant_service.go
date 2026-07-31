package service

import (
	"context"
	"database/sql"
	"errors"

	db "paylater-backend/internal/db"
	"paylater-backend/internal/dto"
	"paylater-backend/internal/auth"
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
func (s *MerchantService) RegisterMerchant(ctx context.Context, req dto.MerchantRegisterRequest) error {

	_, err := s.queries.GetMerchantByEmail(ctx, req.Email)

	if err == nil {
		return errors.New("merchant already exists")
	}

	if err != sql.ErrNoRows {
		return err
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return err
	}

	_, err = s.queries.CreateMerchant(ctx, db.CreateMerchantParams{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: hashedPassword,
	})

	return err
}
//login merchant
func (s *MerchantService) LoginMerchant(ctx context.Context, req dto.MerchantLoginRequest) (string, error) {

	merchant, err := s.queries.GetMerchantByEmail(ctx, req.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = auth.ComparePassword(merchant.Password, req.Password)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := auth.GenerateJWT(merchant.ID, auth.RoleMerchant)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *MerchantService) GetMyProfile(
	ctx context.Context,
	merchantID int32,
) (db.Merchant, error) {

	return s.queries.GetMerchantByID(ctx, merchantID)
}

func (s *MerchantService) UpdateMyProfile(
	ctx context.Context,
	merchantID int32,
	req db.UpdateMerchantParams,
) error {

	req.ID = merchantID

	return s.queries.UpdateMerchant(ctx, req)
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