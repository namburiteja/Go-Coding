package customer

import (
	"time"

	db "paylater/services/customer/internal/db"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// CustomerResponse is the public customer shape — password/hash is never included.
type CustomerResponse struct {
	ID             int32      `json:"id"`
	Name           string     `json:"name"`
	Email          string     `json:"email"`
	CreditLimit    string     `json:"credit_limit"`
	TotalDue       *string    `json:"total_due"`
	PaymentDueDate time.Time  `json:"payment_due_date"`
	Status         *string    `json:"status"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
}

type CreditSnapshotResponse struct {
	ID             int32     `json:"id"`
	CreditLimit    string    `json:"credit_limit"`
	TotalDue       *string   `json:"total_due"`
	PaymentDueDate time.Time `json:"payment_due_date"`
	Status         *string   `json:"status"`
}

type UpdateDueRequest struct {
	TotalDue string `json:"total_due" binding:"required"`
}

func toCustomerResponse(c db.Customer) CustomerResponse {
	resp := CustomerResponse{
		ID:             c.ID,
		Name:           c.Name,
		Email:          c.Email,
		CreditLimit:    c.CreditLimit,
		PaymentDueDate: c.PaymentDueDate,
	}
	if c.TotalDue.Valid {
		resp.TotalDue = &c.TotalDue.String
	}
	if c.Status.Valid {
		status := string(c.Status.CustomersStatus)
		resp.Status = &status
	}
	if c.CreatedAt.Valid {
		t := c.CreatedAt.Time
		resp.CreatedAt = &t
	}
	return resp
}

func toCustomerResponses(customers []db.Customer) []CustomerResponse {
	out := make([]CustomerResponse, 0, len(customers))
	for _, c := range customers {
		out = append(out, toCustomerResponse(c))
	}
	return out
}

// CalculatePaymentDueDate returns the next applicable 5th:
// day < 5 → this month's 5th; day >= 5 → next month's 5th.
func CalculatePaymentDueDate(now time.Time) time.Time {
	year, month, day := now.Date()
	location := now.Location()

	if day < 5 {
		return time.Date(year, month, 5, 0, 0, 0, 0, location)
	}

	return time.Date(year, month+1, 5, 0, 0, 0, 0, location)
}
