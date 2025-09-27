package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from a .env file if it exists
func LoadEnv() {
	// Check if running in a production environment (adjust condition as needed)
	// For example, if GIN_MODE is set to release, we might not want to load .env
	if os.Getenv("GIN_MODE") != "release" {
		// Check if .env file exists in the current working directory
		if _, err := os.Stat(".env"); err == nil {
			// Load the .env file in the current directory
			if err := godotenv.Load(".env"); err != nil {
				log.Printf("Warning: Error loading .env file from current directory: %v", err)
			} else {
				log.Println("Successfully loaded .env file from current directory")
			}
		} else if os.IsNotExist(err) {
			// Check for .env file in the parent directory
			if _, err := os.Stat("../.env"); err == nil {
				if err := godotenv.Load("../.env"); err != nil {
					log.Printf("Warning: Error loading .env file from parent directory: %v", err)
				} else {
					log.Println("Successfully loaded .env file from parent directory")
				}
			} else {
				// No .env file found in current or parent directory
				log.Println("Warning: .env file not found in current or parent directory, using OS environment variables")
			}
		} else {
			// Some other error occurred while checking for the file
			log.Printf("Warning: Error checking for .env file: %v", err)
		}
	} else {
		log.Println("Running in release mode, not loading .env file")
	}
}

// GetEnvOrDefault retrieves an environment variable or returns a default value
func GetEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
