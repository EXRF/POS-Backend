package db

import (
	"fmt"
	"log"
	"os"

	"github.com/EXRF/POS-Backend/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB holds the GORM database connection
var DB *gorm.DB

// ConnectDB establishes a connection to the PostgreSQL database using GORM
func ConnectDB() (*gorm.DB, error) {
	// Load environment variables or use default values
	host := utils.GetEnvOrDefault("DB_HOST", "localhost")
	port := utils.GetEnvOrDefault("DB_PORT", "5432")
	user := utils.GetEnvOrDefault("DB_USER", "postgres")
	password := utils.GetEnvOrDefault("DB_PASSWORD", "postgres")
	dbname := utils.GetEnvOrDefault("DB_NAME", "pos")

	// Connection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Initialize GORM with logging based on environment
	var gormLogger logger.Interface
	if os.Getenv("GIN_MODE") == "release" {
		gormLogger = logger.Default
	} else {
		gormLogger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Test the connection by pinging the database
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	if err = sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL database using GORM")
	DB = db
	return db, nil
}

// CloseDB closes the database connection
func CloseDB() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
