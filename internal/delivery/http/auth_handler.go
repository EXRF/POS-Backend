package http

import (
	"errors"
	"net/http"

	"github.com/EXRF/POS-Backend/internal/entities"
	"github.com/EXRF/POS-Backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// RegisterUser handles user registration
func (h *Handler) RegisterUser(c *gin.Context) {
	var req entities.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		var ve validator.ValidationErrors
		if ok := errors.As(err, &ve); ok {
			fe := ve[0]
			utils.JSONError(c, http.StatusBadRequest, utils.ValidationErrorMessage(fe))
			return
		}

		utils.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.userUsecase.RegisterUser(req)
	if err != nil {
		utils.JSONError(c, utils.MapErrorToStatus(err), err.Error())
		return
	}

	utils.JSONResponse(c, http.StatusCreated, "User registered successfully", nil)
}
