package handler

import (
	"net/http"
	"strconv"

	db "paylater-backend/internal/db"
	"paylater-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	service *service.AdminService
}

func NewAdminHandler(service *service.AdminService) *AdminHandler {
	return &AdminHandler{
		service: service,
	}
}

// POST /admins
func (h *AdminHandler) CreateAdmin(c *gin.Context) {
	var req db.CreateAdminParams

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.service.CreateAdmin(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Admin created successfully",
	})
}

// GET /admins
func (h *AdminHandler) GetAllAdmins(c *gin.Context) {
	admins, err := h.service.GetAllAdmins(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, admins)
}

// GET /admins/:id
func (h *AdminHandler) GetAdminByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Admin ID",
		})
		return
	}

	admin, err := h.service.GetAdminByID(c.Request.Context(), int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Admin not found",
		})
		return
	}

	c.JSON(http.StatusOK, admin)
}

// PUT /admins/:id
func (h *AdminHandler) UpdateAdmin(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Admin ID",
		})
		return
	}

	var req db.UpdateAdminParams

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	req.ID = int32(id)

	err = h.service.UpdateAdmin(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Admin updated successfully",
	})
}

// DELETE /admins/:id
func (h *AdminHandler) DeleteAdminByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Admin ID",
		})
		return
	}

	err = h.service.DeleteAdminByID(c.Request.Context(), int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Admin deleted successfully",
	})
}