package ledger

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// MerchantCommissionHTTP implements MerchantCommissionPort over REST.
type MerchantCommissionHTTP struct {
	baseURL    string
	httpClient *http.Client
}

func NewMerchantCommissionHTTP(baseURL string) *MerchantCommissionHTTP {
	return &MerchantCommissionHTTP{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

type commissionResponse struct {
	ID                   int32   `json:"id"`
	CommissionPercentage *string `json:"commission_percentage"`
}

func (c *MerchantCommissionHTTP) GetCommission(ctx context.Context, merchantID int32) (MerchantCommission, error) {
	url := fmt.Sprintf("%s/internal/merchants/%d/commission", c.baseURL, merchantID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return MerchantCommission{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return MerchantCommission{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return MerchantCommission{}, errors.New("merchant not found")
	}
	if resp.StatusCode != http.StatusOK {
		return MerchantCommission{}, fmt.Errorf("merchant service returned status %d", resp.StatusCode)
	}

	var body commissionResponse
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return MerchantCommission{}, err
	}

	commission := MerchantCommission{ID: body.ID}
	if body.CommissionPercentage != nil {
		commission.CommissionPercentage = sql.NullString{
			String: *body.CommissionPercentage,
			Valid:  true,
		}
	}

	return commission, nil
}
