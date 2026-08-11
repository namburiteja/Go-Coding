package merchant

import (
	"context"
	"database/sql"
	"errors"

	db "paylater/services/merchant/internal/db"
	"paylater/shared/auth"
)

type Service struct {
	queries *db.Queries
}

func NewService(queries *db.Queries) *Service {
	return &Service{queries: queries}
}

func (s *Service) RegisterMerchant(ctx context.Context, req RegisterRequest) error {
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

func (s *Service) LoginMerchant(ctx context.Context, req LoginRequest) (string, error) {
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

func (s *Service) GetMyProfile(ctx context.Context, merchantID int32) (db.Merchant, error) {
	return s.queries.GetMerchantByID(ctx, merchantID)
}

func (s *Service) UpdateMyProfile(ctx context.Context, merchantID int32, req db.UpdateMerchantParams) error {
	req.ID = merchantID
	return s.queries.UpdateMerchant(ctx, req)
}

func (s *Service) GetMerchantByID(ctx context.Context, id int32) (db.Merchant, error) {
	return s.queries.GetMerchantByID(ctx, id)
}

func (s *Service) GetAllMerchants(ctx context.Context) ([]db.Merchant, error) {
	return s.queries.GetAllMerchants(ctx)
}

func (s *Service) UpdateMerchantCommission(ctx context.Context, arg db.UpdateMerchantCommissionParams) error {
	return s.queries.UpdateMerchantCommission(ctx, arg)
}

func (s *Service) UpdateMerchant(ctx context.Context, arg db.UpdateMerchantParams) error {
	return s.queries.UpdateMerchant(ctx, arg)
}

func (s *Service) DeleteMerchant(ctx context.Context, id int32) error {
	return s.queries.DeleteMerchantById(ctx, id)
}

func (s *Service) GetMerchantNames(ctx context.Context) ([]db.GetMerchantNamesRow, error) {
	return s.queries.GetMerchantNames(ctx)
}
