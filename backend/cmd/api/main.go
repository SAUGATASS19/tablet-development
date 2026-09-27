package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"backend/internal/database"
	"backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	//Connect to database
	database.Connect()

	// Initialize the Gin engine with default settings (includes logging and crash recovery)
	r := gin.Default()

	// Create a versioned routing group for the mobile app
	v1 := r.Group("/api/v1")
	{
		// The health check is now accessible at /api/v1/health
		v1.GET("/health", handlers.HealthCheck)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting safely.")
}
