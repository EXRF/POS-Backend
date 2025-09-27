package http

import (
	"github.com/EXRF/POS-Backend/internal/repository"
	"github.com/EXRF/POS-Backend/internal/usecases"
	"gorm.io/gorm"
)

// ProvideHandler wires up dependencies and returns a Handler instance.
func ProvideHandler(db *gorm.DB) *Handler {
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	return NewHandler(userUsecase)
}
