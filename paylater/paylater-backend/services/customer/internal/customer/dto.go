package customer

import "time"

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
