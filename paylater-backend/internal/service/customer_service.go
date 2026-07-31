package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	db "paylater-backend/internal/db"
	"paylater-backend/internal/auth"
	"paylater-backend/internal/dto"
	"paylater-backend/internal/utils"

)

type CustomerService struct {
	queries *db.Queries
}

func NewCustomerService(queries *db.Queries) *CustomerService {
	return &CustomerService{
		queries: queries,
	}
}

func (s *CustomerService) RegisterCustomer(ctx context.Context, req dto.CustomerRegisterRequest) error {

	_, err := s.queries.GetCustomerByEmail(ctx, req.Email)

	if err == nil {
		return errors.New("customer already exists")
	}

	if err != sql.ErrNoRows {
		return err
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return err
	}

	dueDate := utils.CalculatePaymentDueDate(time.Now())

	_, err = s.queries.CreateCustomer(ctx, db.CreateCustomerParams{
		Name:           req.Name,
		Email:          req.Email,
		Password:       hashedPassword,
		PaymentDueDate: dueDate,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *CustomerService) LoginCustomer(ctx context.Context, req dto.CustomerLoginRequest) (string, error) {

	customer, err := s.queries.GetCustomerByEmail(ctx, req.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = auth.ComparePassword(customer.Password, req.Password)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := auth.GenerateJWT(customer.ID, auth.RoleCustomer)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *CustomerService) GetMyProfile(
	ctx context.Context,
	customerID int32,
) (db.Customer, error) {

	return s.queries.GetCustomerByID(ctx, customerID)
}

func (s *CustomerService) UpdateMyProfile(
	ctx context.Context,
	customerID int32,
	req db.UpdateCustomerParams,
) error {

	req.ID = customerID

	return s.queries.UpdateCustomer(ctx, req)
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
