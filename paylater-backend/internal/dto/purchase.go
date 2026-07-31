package dto

type PurchaseRequest struct {
	MerchantID int32  `json:"merchantId" binding:"required"`
	Amount     string `json:"amount" binding:"required"`
}