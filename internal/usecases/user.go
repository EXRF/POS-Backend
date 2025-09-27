package usecases

import (
	"errors"

	"github.com/EXRF/POS-Backend/internal/entities"
	"github.com/EXRF/POS-Backend/internal/repository"
	"github.com/EXRF/POS-Backend/pkg/utils"
)

// UserUsecase defines the interface for user-related business logic
type UserUsecase interface {
	RegisterUser(req entities.CreateUserRequest) (*entities.User, error)
}

// UserUsecaseImpl implements UserUsecase interface
type UserUsecaseImpl struct {
	userRepo repository.UserRepository
}

// NewUserUsecase creates a new instance of UserUsecaseImpl
func NewUserUsecase(userRepo repository.UserRepository) *UserUsecaseImpl {
	return &UserUsecaseImpl{
		userRepo: userRepo,
	}
}

// RegisterUser registers a new user with hashed password
func (uc *UserUsecaseImpl) RegisterUser(req entities.CreateUserRequest) (*entities.User, error) {
	// Check if user with username already exists
	existingUser, _ := uc.userRepo.GetUserByUsername(req.Username)
	if existingUser != nil {
		return nil, errors.New(utils.ErrUserExists)
	}

	// Check if user with email already exists
	existingUser, _ = uc.userRepo.GetUserByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New(utils.ErrEmailExists)
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("Failed to hash password")
	}

	user := &entities.User{
		Username:    req.Username,
		Email:       req.Email,
		Password:    hashedPassword,
		FirstName:   &req.FirstName,
		LastName:    &req.LastName,
		PhoneNumber: &req.PhoneNumber,
		IsActive:    true,
	}

	err = uc.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
