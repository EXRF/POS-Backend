package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/EXRF/POS-Backend/pkg/db"
)

// RunServer starts the server with graceful shutdown
func RunServer() {
	// Connect to database
	pg, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	// Setup router
	r := SetupRouter(pg)
	port := GetPort()

	// Create HTTP server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Block until signal is received
	<-quit
	log.Println("Shutting down server...")

	// Close database connection before shutdown
	if err := db.CloseDB(); err != nil {
		log.Printf("Error closing database connection: %v", err)
	}

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
