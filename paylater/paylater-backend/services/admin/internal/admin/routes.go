package admin

import (
	"paylater/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, h *Handler) {
	admins := router.Group("/admins")
	{
		admins.POST("/register", h.RegisterAdmin)
		admins.POST("/login", h.LoginAdmin)

		admins.GET(
			"",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			h.GetAllAdmins,
		)

		admins.GET(
			"/:id",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			h.GetAdminByID,
		)

		admins.PUT(
			"/:id",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			h.UpdateAdmin,
		)

		admins.DELETE(
			"/:id",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			h.DeleteAdminByID,
		)
	}
}
