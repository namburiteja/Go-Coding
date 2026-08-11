package middleware

import (
	"crypto/subtle"
	"net/http"

	"paylater/shared/internalauth"

	"github.com/gin-gonic/gin"
)

// InternalServiceAuth protects /internal/* routes with INTERNAL_SERVICE_TOKEN.
// It does not accept end-user JWTs — only the shared service token header.
func InternalServiceAuth() gin.HandlerFunc {
	expected := internalauth.Token()

	return func(c *gin.Context) {
		provided := c.GetHeader(internalauth.Header)
		if provided == "" ||
			subtle.ConstantTimeCompare([]byte(provided), []byte(expected)) != 1 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized internal request",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
