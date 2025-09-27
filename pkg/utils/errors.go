package utils

import "net/http"

// Error types
var (
	ErrUserExists   = "Username already exists"
	ErrEmailExists  = "Email already exists"
	ErrUserNotFound = "User not found"
)

// MapErrorToStatus maps known errors to HTTP status codes
func MapErrorToStatus(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err.Error() {
	case ErrUserExists, ErrEmailExists:
		return http.StatusConflict
	case ErrUserNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
