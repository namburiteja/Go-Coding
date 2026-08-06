package ledger

import (
	"time"

	db "paylater/services/ledger/internal/db"
)

type PurchaseRequest struct {
	MerchantID int32  `json:"merchantId" binding:"required"`
	Amount     string `json:"amount" binding:"required"`
}

type PaybackRequest struct {
	Amount string `json:"amount" binding:"required"`
}

// MerchantFeeTotal is returned to Report over the internal HTTP API.
type MerchantFeeTotal struct {
	MerchantID        int32  `json:"merchant_id"`
	TotalFeeCollected string `json:"total_fee_collected"`
}

// TransactionResponse is the public JSON shape for transaction APIs.
type TransactionResponse struct {
	ID                   int32      `json:"id"`
	CustomerID           int32      `json:"customer_id"`
	MerchantID           *int32     `json:"merchant_id"`
	TransactionType      string     `json:"transaction_type"`
	Amount               string     `json:"amount"`
	CommissionPercentage *string    `json:"commission_percentage"`
	CommissionAmount     *string    `json:"commission_amount"`
	TransactionDate      *time.Time `json:"transaction_date"`
}

func toTransactionResponse(t db.Transaction) TransactionResponse {
	resp := TransactionResponse{
		ID:              t.ID,
		CustomerID:      t.CustomerID,
		TransactionType: string(t.TransactionType),
		Amount:          t.Amount,
	}
	if t.MerchantID.Valid {
		id := t.MerchantID.Int32
		resp.MerchantID = &id
	}
	if t.CommissionPercentage.Valid {
		v := t.CommissionPercentage.String
		resp.CommissionPercentage = &v
	}
	if t.CommissionAmount.Valid {
		v := t.CommissionAmount.String
		resp.CommissionAmount = &v
	}
	if t.TransactionDate.Valid {
		v := t.TransactionDate.Time
		resp.TransactionDate = &v
	}
	return resp
}

func toTransactionResponses(items []db.Transaction) []TransactionResponse {
	out := make([]TransactionResponse, 0, len(items))
	for _, item := range items {
		out = append(out, toTransactionResponse(item))
	}
	return out
}
