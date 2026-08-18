package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const allowedOrigin = "http://localhost:5173"

// CORS allows the Vite React app to call the gateway.
// OPTIONS preflight is answered here so it never falls through to NoRoute/proxy 404s.
// Credentials are not enabled: the frontend sends JWT via Authorization, not cookies.
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin == allowedOrigin {
			c.Header("Access-Control-Allow-Origin", allowedOrigin)
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Header("Vary", "Origin")
		}

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
