package routes

import (
	"paylater/shared/proxy"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures the API edge: every domain path is reverse-proxied
// to its owning service. No domain handlers remain on the gateway.
func SetupRoutes(
	router *gin.Engine,
	adminServiceURL string,
	merchantServiceURL string,
	customerServiceURL string,
	ledgerServiceURL string,
	reportServiceURL string,
) {
	ledgerProxy := proxy.Forward(ledgerServiceURL)
	router.POST("/customers/purchase", ledgerProxy)
	router.POST("/customers/payback", ledgerProxy)
	router.GET("/customers/me/transactions", ledgerProxy)
	router.GET("/merchants/me/transactions", ledgerProxy)
	router.GET("/transactions", ledgerProxy)

	customerProxy := proxy.Forward(customerServiceURL)
	router.Any("/customers", customerProxy)
	router.Any("/customers/*path", customerProxy)

	merchantProxy := proxy.Forward(merchantServiceURL)
	router.Any("/merchants", merchantProxy)
	router.Any("/merchants/*path", merchantProxy)

	reportProxy := proxy.Forward(reportServiceURL)
	router.Any("/reports", reportProxy)
	router.Any("/reports/*path", reportProxy)

	adminProxy := proxy.Forward(adminServiceURL)
	router.Any("/admins", adminProxy)
	router.Any("/admins/*path", adminProxy)
}
