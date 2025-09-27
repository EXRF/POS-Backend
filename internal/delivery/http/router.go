package http

import (
	"os"

	"github.com/EXRF/POS-Backend/pkg/middlewares"
	"github.com/EXRF/POS-Backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes configures all the API routes
func SetupRoutes(router *gin.Engine, handler *Handler) {
	// Health check endpoint
	router.GET("/health", handler.HealthCheck)

	// Group routes under /api/v1
	v1 := router.Group("/api/v1")
	{
		// Auth routes
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", handler.RegisterUser)
		}
	}
}

// SetupRouter initializes the router with database connection, repositories, usecases, and handlers
func SetupRouter(db *gorm.DB) *gin.Engine {
	// Set Gin mode based on environment
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()

	// Add custom logger middleware
	r.Use(middlewares.Logger())

	// Add recovery middleware to handle panics
	r.Use(middlewares.RecoveryWithLogger())

	// Add error handling middleware
	r.Use(middlewares.ErrorHandler())

	handler := ProvideHandler(db)
	SetupRoutes(r, handler)

	return r
}

// GetPort returns the port to run the server on
func GetPort() string {
	return utils.GetEnvOrDefault("PORT", "8080")
}
