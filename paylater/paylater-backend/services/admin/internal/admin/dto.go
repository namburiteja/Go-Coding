package admin

import (
	"time"

	db "paylater/services/admin/internal/db"
	"paylater/shared/auth"
)

// RegisterRequest creates another ADMIN. Role is never accepted from the client.
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

// AdminResponse is the public admin shape — password/hash is never included.
type AdminResponse struct {
	ID        int32      `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Role      string     `json:"role"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

func toAdminResponse(a db.Admin) AdminResponse {
	resp := AdminResponse{
		ID:    a.ID,
		Name:  a.Name,
		Email: a.Email,
		Role:  auth.RoleAdmin,
	}
	if a.CreatedAt.Valid {
		t := a.CreatedAt.Time
		resp.CreatedAt = &t
	}
	return resp
}

func toAdminResponses(admins []db.Admin) []AdminResponse {
	out := make([]AdminResponse, 0, len(admins))
	for _, a := range admins {
		out = append(out, toAdminResponse(a))
	}
	return out
}
