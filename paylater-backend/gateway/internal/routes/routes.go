package routes

import (
	"net/http"
	"strings"

	"paylater/shared/proxy"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures the API edge as a pure reverse proxy.
// A single catch-all avoids Gin conflicts between static and wildcard routes.
func SetupRoutes(
	router *gin.Engine,
	adminServiceURL string,
	merchantServiceURL string,
	customerServiceURL string,
	ledgerServiceURL string,
	reportServiceURL string,
) {
	ledgerProxy := proxy.Forward(ledgerServiceURL)
	customerProxy := proxy.Forward(customerServiceURL)
	merchantProxy := proxy.Forward(merchantServiceURL)
	reportProxy := proxy.Forward(reportServiceURL)
	adminProxy := proxy.Forward(adminServiceURL)

	dispatch := func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method

		switch {
		case isLedgerRoute(method, path):
			ledgerProxy(c)
		case path == "/customers" || strings.HasPrefix(path, "/customers/"):
			customerProxy(c)
		case path == "/merchants" || strings.HasPrefix(path, "/merchants/"):
			merchantProxy(c)
		case path == "/reports" || strings.HasPrefix(path, "/reports/"):
			reportProxy(c)
		case path == "/admins" || strings.HasPrefix(path, "/admins/"):
			adminProxy(c)
		default:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "not found"})
		}
	}

	// No overlapping Gin route trees — dispatch everything through NoRoute.
	router.NoRoute(dispatch)
	router.NoMethod(dispatch)
}

func isLedgerRoute(method, path string) bool {
	switch {
	case method == http.MethodPost && path == "/customers/purchase":
		return true
	case method == http.MethodPost && path == "/customers/payback":
		return true
	case method == http.MethodGet && path == "/customers/me/transactions":
		return true
	case method == http.MethodGet && path == "/merchants/me/transactions":
		return true
	case method == http.MethodGet && path == "/transactions":
		return true
	default:
		return false
	}
}
