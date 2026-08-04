package ledger

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
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
			Timeout: 5 * time.Second,
		},
	}
}

type creditSnapshotResponse struct {
	ID             int32     `json:"id"`
	CreditLimit    string    `json:"credit_limit"`
	TotalDue       *string   `json:"total_due"`
	PaymentDueDate time.Time `json:"payment_due_date"`
	Status         *string   `json:"status"`
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
		return CreditAccount{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return CreditAccount{}, errors.New("customer not found")
	}
	if resp.StatusCode != http.StatusOK {
		return CreditAccount{}, fmt.Errorf("customer service returned status %d", resp.StatusCode)
	}

	var body creditSnapshotResponse
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return CreditAccount{}, err
	}

	account := CreditAccount{
		ID:             body.ID,
		CreditLimit:    body.CreditLimit,
		PaymentDueDate: body.PaymentDueDate,
	}
	if body.TotalDue != nil {
		account.TotalDue = sql.NullString{String: *body.TotalDue, Valid: true}
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
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("customer service returned status %d", resp.StatusCode)
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
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("customer service returned status %d", resp.StatusCode)
	}
	return nil
}
