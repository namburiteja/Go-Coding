package dto

type PaybackRequest struct {
	Amount string `json:"amount" binding:"required"`
}