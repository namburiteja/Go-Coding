package ledger

type PurchaseRequest struct {
	MerchantID int32  `json:"merchantId" binding:"required"`
	Amount     string `json:"amount" binding:"required"`
}

type PaybackRequest struct {
	Amount string `json:"amount" binding:"required"`
}
