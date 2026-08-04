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

func CalculatePaymentDueDate(now time.Time) time.Time {
	year, month, day := now.Date()
	location := now.Location()

	if day <= 5 {
		return time.Date(year, month, 5, 0, 0, 0, 0, location)
	}

	return time.Date(year, month+1, 5, 0, 0, 0, 0, location)
}
