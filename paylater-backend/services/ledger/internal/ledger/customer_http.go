package ledger

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// CustomerCreditHTTP implements CustomerCreditPort over REST.
type CustomerCreditHTTP struct {
	baseURL    string
	httpClient *http.Client
}

func NewCustomerCreditHTTP(baseURL string) *CustomerCreditHTTP {
	return &CustomerCreditHTTP{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type creditSnapshotResponse struct {
	ID             int32           `json:"id"`
	CreditLimit    json.RawMessage `json:"credit_limit"`
	TotalDue       json.RawMessage `json:"total_due"`
	PaymentDueDate time.Time       `json:"payment_due_date"`
	Status         *string         `json:"status"`
}

type updateDueRequest struct {
	TotalDue string `json:"total_due"`
}

func (c *CustomerCreditHTTP) GetForUpdate(ctx context.Context, customerID int32) (CreditAccount, error) {
	url := fmt.Sprintf("%s/internal/customers/%d/credit", c.baseURL, customerID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return CreditAccount{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return CreditAccount{}, fmt.Errorf("customer service unreachable: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return CreditAccount{}, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return CreditAccount{}, errors.New("customer not found")
	}
	if resp.StatusCode != http.StatusOK {
		return CreditAccount{}, fmt.Errorf("customer service returned status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var body creditSnapshotResponse
	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		return CreditAccount{}, fmt.Errorf("decode customer credit: %w", err)
	}

	creditLimit, err := rawToString(body.CreditLimit)
	if err != nil {
		return CreditAccount{}, fmt.Errorf("invalid credit_limit: %w", err)
	}

	account := CreditAccount{
		ID:             body.ID,
		CreditLimit:    creditLimit,
		PaymentDueDate: body.PaymentDueDate,
	}

	if len(body.TotalDue) > 0 && string(body.TotalDue) != "null" {
		due, err := rawToString(body.TotalDue)
		if err != nil {
			return CreditAccount{}, fmt.Errorf("invalid total_due: %w", err)
		}
		account.TotalDue = sql.NullString{String: due, Valid: true}
	}

	if body.Status != nil {
		account.Status = *body.Status
		account.StatusValid = true
	}

	return account, nil
}

func (c *CustomerCreditHTTP) UpdateDue(ctx context.Context, customerID int32, totalDue string) error {
	url := fmt.Sprintf("%s/internal/customers/%d/due", c.baseURL, customerID)

	payload, err := json.Marshal(updateDueRequest{TotalDue: totalDue})
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("customer service unreachable: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("customer service returned status %d: %s", resp.StatusCode, string(bodyBytes))
	}
	return nil
}

func (c *CustomerCreditHTTP) Block(ctx context.Context, customerID int32) error {
	url := fmt.Sprintf("%s/internal/customers/%d/block", c.baseURL, customerID)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, nil)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("customer service unreachable: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("customer service returned status %d: %s", resp.StatusCode, string(bodyBytes))
	}
	return nil
}

func rawToString(raw json.RawMessage) (string, error) {
	if len(raw) == 0 || string(raw) == "null" {
		return "", nil
	}
	var asString string
	if err := json.Unmarshal(raw, &asString); err == nil {
		return asString, nil
	}
	var asNumber json.Number
	if err := json.Unmarshal(raw, &asNumber); err == nil {
		return asNumber.String(), nil
	}
	return "", fmt.Errorf("unsupported value %s", string(raw))
}
