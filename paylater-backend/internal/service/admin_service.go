package service

import (
	"context"
	"errors"
	"database/sql"

	db "paylater-backend/internal/db"
	"paylater-backend/internal/auth"
	"paylater-backend/internal/dto"
)

type AdminService struct {
	queries *db.Queries
}

func NewAdminService(queries *db.Queries) *AdminService {
	return &AdminService{
		queries: queries,
	}
}

func (s *AdminService) RegisterAdmin(ctx context.Context, req dto.AdminRegisterRequest) error {

	// Check if email already exists
	_, err := s.queries.GetAdminByEmail(ctx, req.Email)

	if err == nil {
		return errors.New("admin already exists with this email")
	}

	if err != sql.ErrNoRows {
		return err
	}

	// Hash password
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return err
	}

	// Save admin
	_, err = s.queries.CreateAdmin(ctx, db.CreateAdminParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	})

	return err
}

func (s *AdminService) LoginAdmin(ctx context.Context, req dto.AdminLoginRequest) (string, error) {

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

// Read by ID
func (s *AdminService) GetAdminByID(ctx context.Context, id int32) (db.Admin, error) {
	return s.queries.GetAdminByID(ctx, id)
}

// Read by Email (used for Login)
func (s *AdminService) GetAdminByEmail(ctx context.Context, email string) (db.Admin, error) {
	return s.queries.GetAdminByEmail(ctx, email)
}

// Read All
func (s *AdminService) GetAllAdmins(ctx context.Context) ([]db.Admin, error) {
	return s.queries.GetAllAdmins(ctx)
}

// Update
func (s *AdminService) UpdateAdmin(ctx context.Context, arg db.UpdateAdminParams) error {
	return s.queries.UpdateAdmin(ctx, arg)
}

// Delete
func (s *AdminService) DeleteAdminByID(ctx context.Context, id int32) error {
	return s.queries.DeleteAdminByID(ctx, id)
}