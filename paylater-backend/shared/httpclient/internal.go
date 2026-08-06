package httpclient

import (
	"context"
	"io"
	"net/http"
	"time"

	"paylater/shared/internalauth"
)

// Internal is a reusable HTTP client for service-to-service calls.
// Every request automatically carries INTERNAL_SERVICE_TOKEN.
// It contains no domain-specific logic — any microservice can use it.
type Internal struct {
	httpClient *http.Client
	token      string
}

// NewInternal builds an Internal client with the given timeout.
// INTERNAL_SERVICE_TOKEN must already be loaded into the environment.
func NewInternal(timeout time.Duration) *Internal {
	return &Internal{
		httpClient: &http.Client{Timeout: timeout},
		token:      internalauth.Token(),
	}
}

// NewRequest creates an HTTP request with context (token is attached on Do).
func (c *Internal) NewRequest(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, method, url, body)
}

// Do sends req after setting the internal service token header.
func (c *Internal) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set(internalauth.Header, c.token)
	return c.httpClient.Do(req)
}
