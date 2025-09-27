package http

import (
	"net/http"

	"github.com/EXRF/POS-Backend/internal/entities"
	"github.com/gin-gonic/gin"
)

// RegisterUser handles user registration
func (h *Handler) RegisterUser(c *gin.Context) {
	var req entities.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.userUsecase.RegisterUser(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}
