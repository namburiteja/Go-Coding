package customer

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	db "paylater/services/customer/internal/db"
	"paylater/shared/auth"
)

type Service struct {
	db      *sql.DB
	queries *db.Queries
}

func NewService(database *sql.DB, queries *db.Queries) *Service {
	return &Service{db: database, queries: queries}
}

func (s *Service) RegisterCustomer(ctx context.Context, req RegisterRequest) error {
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

	dueDate := CalculatePaymentDueDate(time.Now())

	_, err = s.queries.CreateCustomer(ctx, db.CreateCustomerParams{
		Name:           req.Name,
		Email:          req.Email,
		Password:       hashedPassword,
		PaymentDueDate: dueDate,
	})
	return err
}

func (s *Service) LoginCustomer(ctx context.Context, req LoginRequest) (string, error) {
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

func (s *Service) GetMyProfile(ctx context.Context, customerID int32) (db.Customer, error) {
	return s.queries.GetCustomerByID(ctx, customerID)
}

func (s *Service) UpdateMyProfile(ctx context.Context, customerID int32, req db.UpdateCustomerParams) error {
	req.ID = customerID
	return s.queries.UpdateCustomer(ctx, req)
}

func (s *Service) GetCustomerByID(ctx context.Context, id int32) (db.Customer, error) {
	return s.queries.GetCustomerByID(ctx, id)
}

func (s *Service) GetAllCustomers(ctx context.Context) ([]db.Customer, error) {
	return s.queries.GetAllCustomers(ctx)
}

func (s *Service) UpdateCustomer(ctx context.Context, arg db.UpdateCustomerParams) error {
	return s.queries.UpdateCustomer(ctx, arg)
}

func (s *Service) DeleteCustomer(ctx context.Context, id int32) error {
	return s.queries.DeleteCustomerById(ctx, id)
}

func (s *Service) GetCreditForUpdate(ctx context.Context, customerID int32) (db.Customer, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return db.Customer{}, err
	}
	defer tx.Rollback()

	customer, err := s.queries.WithTx(tx).GetCustomerByIDForUpdate(ctx, customerID)
	if err != nil {
		return db.Customer{}, err
	}
	if err := tx.Commit(); err != nil {
		return db.Customer{}, err
	}
	return customer, nil
}

func (s *Service) UpdateDue(ctx context.Context, customerID int32, totalDue string) error {
	newDue, err := strconv.ParseFloat(totalDue, 64)
	if err != nil {
		return errors.New("invalid total_due")
	}

	customer, err := s.queries.GetCustomerByID(ctx, customerID)
	if err != nil {
		return err
	}

	currentDue := 0.0
	if customer.TotalDue.Valid {
		currentDue, err = strconv.ParseFloat(customer.TotalDue.String, 64)
		if err != nil {
			return err
		}
	}

	paymentDueDate := customer.PaymentDueDate
	status := db.NullCustomersStatus{
		CustomersStatus: db.CustomersStatusACTIVE,
		Valid:           true,
	}
	if customer.Status.Valid {
		status = customer.Status
	}

	now := time.Now()
	switch {
	case newDue == 0:
		paymentDueDate = CalculatePaymentDueDate(now)
		status = db.NullCustomersStatus{
			CustomersStatus: db.CustomersStatusACTIVE,
			Valid:           true,
		}
	case currentDue == 0 && newDue > 0:
		paymentDueDate = CalculatePaymentDueDate(now)
	}

	return s.queries.UpdateCustomerCreditState(ctx, db.UpdateCustomerCreditStateParams{
		ID: customerID,
		TotalDue: sql.NullString{
			String: totalDue,
			Valid:  true,
		},
		PaymentDueDate: paymentDueDate,
		Status:         status,
	})
}

func (s *Service) BlockCustomer(ctx context.Context, customerID int32) error {
	return s.queries.UpdateCustomerStatus(ctx, db.UpdateCustomerStatusParams{
		ID: customerID,
		Status: db.NullCustomersStatus{
			CustomersStatus: db.CustomersStatusBLOCKED,
			Valid:           true,
		},
	})
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
