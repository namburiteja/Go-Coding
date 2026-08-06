package ledger

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
			Timeout: 10 * time.Second,
		},
	}
}

type commissionResponse struct {
	ID                   int32           `json:"id"`
	CommissionPercentage json.RawMessage `json:"commission_percentage"`
}

func (c *MerchantCommissionHTTP) GetCommission(ctx context.Context, merchantID int32) (MerchantCommission, error) {
	url := fmt.Sprintf("%s/internal/merchants/%d/commission", c.baseURL, merchantID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return MerchantCommission{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return MerchantCommission{}, fmt.Errorf("merchant service unreachable: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return MerchantCommission{}, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return MerchantCommission{}, errors.New("merchant not found")
	}
	if resp.StatusCode != http.StatusOK {
		return MerchantCommission{}, fmt.Errorf("merchant service returned status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var body commissionResponse
	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		return MerchantCommission{}, fmt.Errorf("decode merchant commission: %w", err)
	}

	commission := MerchantCommission{ID: body.ID}
	if len(body.CommissionPercentage) > 0 && string(body.CommissionPercentage) != "null" {
		percentage, err := rawToString(body.CommissionPercentage)
		if err != nil {
			return MerchantCommission{}, fmt.Errorf("invalid commission_percentage: %w", err)
		}
		commission.CommissionPercentage = sql.NullString{
			String: percentage,
			Valid:  true,
		}
	}

	return commission, nil
}
