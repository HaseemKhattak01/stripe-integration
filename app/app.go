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
	authService := services.NewStripeService(cfg.StripeKey)
	validationService := validation.NewValidationService()

	router := setupRouter(authService, validationService)

	server := startServer(router)

	gracefulShutdown(server)
}

func setupRouter(authService *services.StripeService, validationService *validation.ValidationService) *gin.Engine {
	router := routes.NewRouter(authService, validationService).Engine
	router.Use(enableCors())
	return router
}

func enableCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func startServer(router *gin.Engine) *http.Server {
	server := &http.Server{
		Addr:    ":8080", 
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return server
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
