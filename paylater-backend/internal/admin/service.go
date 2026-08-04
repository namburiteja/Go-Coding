package admin

import (
	"context"
	"database/sql"
	"errors"

	db "paylater-backend/internal/admin/db"
	"paylater-backend/internal/platform/auth"
)

type Service struct {
	queries *db.Queries
}

func NewService(queries *db.Queries) *Service {
	return &Service{queries: queries}
}

func (s *Service) RegisterAdmin(ctx context.Context, req RegisterRequest) error {
	_, err := s.queries.GetAdminByEmail(ctx, req.Email)
	if err == nil {
		return errors.New("admin already exists with this email")
	}
	if err != sql.ErrNoRows {
		return err
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return err
	}

	_, err = s.queries.CreateAdmin(ctx, db.CreateAdminParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	})
	return err
}

func (s *Service) LoginAdmin(ctx context.Context, req LoginRequest) (string, error) {
	admin, err := s.queries.GetAdminByEmail(ctx, req.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = auth.ComparePassword(admin.Password, req.Password)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := auth.GenerateJWT(admin.ID, auth.RoleAdmin)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) GetAdminByID(ctx context.Context, id int32) (db.Admin, error) {
	return s.queries.GetAdminByID(ctx, id)
}

func (s *Service) GetAdminByEmail(ctx context.Context, email string) (db.Admin, error) {
	return s.queries.GetAdminByEmail(ctx, email)
}

func (s *Service) GetAllAdmins(ctx context.Context) ([]db.Admin, error) {
	return s.queries.GetAllAdmins(ctx)
}

func (s *Service) UpdateAdmin(ctx context.Context, arg db.UpdateAdminParams) error {
	return s.queries.UpdateAdmin(ctx, arg)
}

func (s *Service) DeleteAdminByID(ctx context.Context, id int32) error {
	return s.queries.DeleteAdminByID(ctx, id)
}
