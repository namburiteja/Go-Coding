package middleware

import (
	"net/http"
	"strings"

	"paylater/shared/config"

	"github.com/gin-gonic/gin"
)

// defaultOrigins cover local Vite (5173) and the Dockerized SPA (8081).
var defaultOrigins = []string{
	"http://localhost:5173",
	"http://localhost:8081",
}

// CORS allows browser apps to call the gateway.
// Origins come from CORS_ALLOWED_ORIGINS (comma-separated). When unset/empty,
// defaults are localhost:5173 and localhost:8081.
// OPTIONS preflight is answered here so it never falls through to NoRoute/proxy 404s.
// Credentials are not enabled: the frontend sends JWT via Authorization, not cookies.
func CORS() gin.HandlerFunc {
	allowed := parseAllowedOrigins(config.GetEnv("CORS_ALLOWED_ORIGINS"))

	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin != "" {
			if _, ok := allowed[origin]; ok {
				c.Header("Access-Control-Allow-Origin", origin)
				c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
				c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
				c.Header("Vary", "Origin")
			}
		}

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func parseAllowedOrigins(raw string) map[string]struct{} {
	out := make(map[string]struct{})
	raw = strings.TrimSpace(raw)
	if raw == "" {
		for _, o := range defaultOrigins {
			out[o] = struct{}{}
		}
		return out
	}

	for _, part := range strings.Split(raw, ",") {
		o := strings.TrimSpace(part)
		if o != "" {
			out[o] = struct{}{}
		}
	}
	if len(out) == 0 {
		for _, o := range defaultOrigins {
			out[o] = struct{}{}
		}
	}
	return out
}
