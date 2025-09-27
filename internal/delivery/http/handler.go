package http

import (
	"net/http"

	"github.com/EXRF/POS-Backend/internal/usecases"
	"github.com/gin-gonic/gin"
)

// Handler holds all the use cases
type Handler struct {
	userUsecase usecases.UserUsecase
}

// NewHandler creates a new Handler instance
func NewHandler(userUsecase usecases.UserUsecase) *Handler {
	return &Handler{
		userUsecase: userUsecase,
	}
}

// HealthCheck returns a simple health check endpoint
func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "POS Backend is running",
	})
}
