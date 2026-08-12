package admin

import (
	"paylater/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, h *Handler) {
	admins := router.Group("/admins")
	{
		// Public: admin login only. Registration requires an existing ADMIN JWT.
		admins.POST("/login", h.LoginAdmin)

		admins.POST(
			"/register",
			middleware.AuthMiddleware(),
			middleware.AdminOnly(),
			h.RegisterAdmin,
		)

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
