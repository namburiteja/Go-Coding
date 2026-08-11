package report

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"paylater/shared/httpclient"
)

// ErrNotFound is returned when an upstream domain service responds 404.
var ErrNotFound = errors.New("not found")

// CustomersAPI fetches customer report data from Customer service.
type CustomersAPI struct {
	baseURL string
	client  *httpclient.Internal
}

func NewCustomersAPI(baseURL string) *CustomersAPI {
	return &CustomersAPI{
		baseURL: baseURL,
		client:  httpclient.NewInternal(10 * time.Second),
	}
}

func (c *CustomersAPI) GetUsersAtCreditLimit(ctx context.Context) (json.RawMessage, error) {
	return c.get(ctx, "/internal/customers/reports/at-credit-limit")
}

func (c *CustomersAPI) GetCustomersWithDue(ctx context.Context) (json.RawMessage, error) {
	return c.get(ctx, "/internal/customers/reports/with-due")
}

func (c *CustomersAPI) GetCustomerDueByName(ctx context.Context, name string) (json.RawMessage, error) {
	path := "/internal/customers/reports/due-by-name/" + url.PathEscape(name)
	return c.get(ctx, path)
}

func (c *CustomersAPI) get(ctx context.Context, path string) (json.RawMessage, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, c.baseURL+path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("customer service unreachable: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrNotFound
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("customer service returned status %d: %s", resp.StatusCode, string(body))
	}
	return json.RawMessage(body), nil
}

// MerchantsAPI fetches merchant name snapshots from Merchant service.
type MerchantsAPI struct {
	baseURL string
	client  *httpclient.Internal
}

func NewMerchantsAPI(baseURL string) *MerchantsAPI {
	return &MerchantsAPI{
		baseURL: baseURL,
		client:  httpclient.NewInternal(10 * time.Second),
	}
}

type merchantName struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (m *MerchantsAPI) GetNames(ctx context.Context) ([]merchantName, error) {
	req, err := m.client.NewRequest(ctx, http.MethodGet, m.baseURL+"/internal/merchants/reports/names", nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("merchant service unreachable: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("merchant service returned status %d: %s", resp.StatusCode, string(body))
	}

	var names []merchantName
	if err := json.Unmarshal(body, &names); err != nil {
		return nil, fmt.Errorf("decode merchant names: %w", err)
	}
	if names == nil {
		names = []merchantName{}
	}
	return names, nil
}

// LedgerAPI fetches fee aggregates from Ledger service.
type LedgerAPI struct {
	baseURL string
	client  *httpclient.Internal
}

func NewLedgerAPI(baseURL string) *LedgerAPI {
	return &LedgerAPI{
		baseURL: baseURL,
		client:  httpclient.NewInternal(10 * time.Second),
	}
}

type merchantFeeTotal struct {
	MerchantID        int32  `json:"merchant_id"`
	TotalFeeCollected string `json:"total_fee_collected"`
}

func (l *LedgerAPI) GetMerchantFeeTotals(ctx context.Context) ([]merchantFeeTotal, error) {
	req, err := l.client.NewRequest(ctx, http.MethodGet, l.baseURL+"/internal/transactions/reports/merchant-fees", nil)
	if err != nil {
		return nil, err
	}

	resp, err := l.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ledger service unreachable: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ledger service returned status %d: %s", resp.StatusCode, string(body))
	}

	var fees []merchantFeeTotal
	if err := json.Unmarshal(body, &fees); err != nil {
		return nil, fmt.Errorf("decode merchant fees: %w", err)
	}
	if fees == nil {
		fees = []merchantFeeTotal{}
	}
	return fees, nil
}
