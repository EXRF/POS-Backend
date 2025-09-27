package repository

import (
	"errors"
	"fmt"

	"github.com/EXRF/POS-Backend/internal/entities"
	"github.com/EXRF/POS-Backend/pkg/utils"
	"gorm.io/gorm"
)

// UserRepository defines the interface for user-related database operations
type UserRepository interface {
	CreateUser(user *entities.User) error
	GetUserByEmail(email string) (*entities.User, error)
	GetUserByUsername(username string) (*entities.User, error)
}

// PostgresUserRepository implements UserRepository interface using PostgreSQL
type userRepository struct {
	db *gorm.DB
}

// NewPostgresUserRepository creates a new instance of PostgresUserRepository
func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db,
	}
}

func (r *userRepository) CreateUser(user *entities.User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		return fmt.Errorf("failed to create user: %w", result.Error)
	}
	return nil
}

func (r *userRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	result := r.db.Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New(utils.ErrUserNotFound)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) GetUserByUsername(username string) (*entities.User, error) {
	var user entities.User
	result := r.db.Where("username = ?", username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New(utils.ErrUserNotFound)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
