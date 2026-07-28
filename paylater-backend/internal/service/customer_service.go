package service

import (
	"context"

	db "paylater-backend/internal/db"
)

type CustomerService struct {
	queries *db.Queries
}

func NewCustomerService(queries *db.Queries) *CustomerService {
	return &CustomerService{
		queries: queries,
	}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, arg db.CreateCustomerParams) error {

	// Business logic will come here later
	// Example:
	// 1. Check email already exists
	// 2. Validate credit limit
	// 3. Other business validations

	_, err := s.queries.CreateCustomer(ctx, arg)
	if err != nil {
		return err
	}

	return nil
}

// Get Customer By ID
func (s *CustomerService) GetCustomerByID(ctx context.Context, id int32) (db.Customer, error) {

	return s.queries.GetCustomerByID(ctx, id)

}

// Get All Customers
func (s *CustomerService) GetAllCustomers(ctx context.Context) ([]db.Customer, error) {

	return s.queries.GetAllCustomers(ctx)

}
// Update Customer
func (s *CustomerService) UpdateCustomer(ctx context.Context, arg db.UpdateCustomerParams) error {

	err := s.queries.UpdateCustomer(ctx, arg)
	if err != nil {
		return err
	}

	return nil
}

// Delete Customer
func (s *CustomerService) DeleteCustomer(ctx context.Context, id int32) error {

	err := s.queries.DeleteCustomerById(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
