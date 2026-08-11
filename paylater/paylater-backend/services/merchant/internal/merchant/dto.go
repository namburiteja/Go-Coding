package merchant

import (
	"time"

	db "paylater/services/merchant/internal/db"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// MerchantResponse is the public merchant shape — password/hash is never included.
type MerchantResponse struct {
	ID                   int32      `json:"id"`
	Name                 string     `json:"name"`
	Email                string     `json:"email"`
	Phone                string     `json:"phone"`
	CommissionPercentage *string    `json:"commission_percentage"`
	CreatedAt            *time.Time `json:"created_at,omitempty"`
}

// CommissionResponse is returned to Ledger over the internal HTTP API.
type CommissionResponse struct {
	ID                   int32   `json:"id"`
	CommissionPercentage *string `json:"commission_percentage"`
}

// MerchantNameResponse is returned to Report over the internal HTTP API.
type MerchantNameResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

// UpdateCommissionRequest is the public body for PUT /merchants/:id/commission.
type UpdateCommissionRequest struct {
	CommissionPercentage string `json:"commission_percentage" binding:"required"`
}

func toMerchantResponse(m db.Merchant) MerchantResponse {
	resp := MerchantResponse{
		ID:    m.ID,
		Name:  m.Name,
		Email: m.Email,
		Phone: m.Phone,
	}
	if m.CommissionPercentage.Valid {
		resp.CommissionPercentage = &m.CommissionPercentage.String
	}
	if m.CreatedAt.Valid {
		t := m.CreatedAt.Time
		resp.CreatedAt = &t
	}
	return resp
}

func toMerchantResponses(merchants []db.Merchant) []MerchantResponse {
	out := make([]MerchantResponse, 0, len(merchants))
	for _, m := range merchants {
		out = append(out, toMerchantResponse(m))
	}
	return out
}
