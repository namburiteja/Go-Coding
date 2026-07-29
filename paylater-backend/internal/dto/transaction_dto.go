package dto

type CreateTransactionRequest struct {
	CustomerID      int32  `json:"customer_id" binding:"required"`
	MerchantID      *int32 `json:"merchant_id"`
	TransactionType string `json:"transaction_type" binding:"required"`
	Amount          string `json:"amount" binding:"required"`
}