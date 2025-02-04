package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HaseemKhattak01/stripe-integration/config"
	"github.com/HaseemKhattak01/stripe-integration/routes"
	"github.com/HaseemKhattak01/stripe-integration/services"
	"github.com/HaseemKhattak01/stripe-integration/validation"
	"github.com/gin-gonic/gin"
)

func StartServer(cfg *config.Config) {
	// Initialize services
	authService := services.NewStripeService(cfg.StripeKey)
	validationService := validation.NewValidationService()

	// Create router
	router := setupRouter(*authService, *validationService)

	// Start server
	server := &http.Server{
		Addr:    ":8080", // Default port set to 8080
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Graceful shutdown
	gracefulShutdown(server)
}

func setupRouter(authService services.StripeService, validationService validation.ValidationService) *gin.Engine {
	r := routes.NewRouter(&authService, &validationService).Engine
	r.Use(enableCors())
	r.Static("/public", "./public")
	return r
}

func enableCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func gracefulShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	log.Println("Server exited properly")
}
