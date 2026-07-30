package service

import (
	"context"

	db "paylater-backend/internal/db"
)

type AdminService struct {
	queries *db.Queries
}

func NewAdminService(queries *db.Queries) *AdminService {
	return &AdminService{
		queries: queries,
	}
}

// Create
func (s *AdminService) CreateAdmin(ctx context.Context, arg db.CreateAdminParams) error {
	_, err := s.queries.CreateAdmin(ctx, arg)
	return err
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