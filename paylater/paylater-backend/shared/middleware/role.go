package middleware

import (
	"net/http"

	"paylater/shared/auth"

	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")

		if role != auth.RoleAdmin {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Access denied. Admins only.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func MerchantOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")

		if role != auth.RoleMerchant {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Access denied. Merchants only.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func CustomerOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")

		if role != auth.RoleCustomer {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Access denied. Customers only.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
