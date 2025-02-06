package controllers

import (
	"log"
	"net/http"

	"github.com/HaseemKhattak01/stripe-integration/models"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	GenerateToken(cardDetails models.CardDetails) (string, error)
}

type ValidationService interface {
	ValidateCardDetails(cardDetails models.CardDetails) error
}

func GenerateToken(c *gin.Context, authService AuthService, validationService ValidationService) {
	var cardDetails models.CardDetails
	if err := c.ShouldBindJSON(&cardDetails); err != nil {
		handleError(c, http.StatusBadRequest, "Error binding JSON", err)
		log.Printf("Error binding JSON: %v", err)
		return
	}

	if err := validationService.ValidateCardDetails(cardDetails); err != nil {
		handleError(c, http.StatusBadRequest, "Validation error", err)
		log.Printf("Validation error: %v", err)
		return
	}

	token, err := authService.GenerateToken(cardDetails)
	if err != nil {
		handleError(c, http.StatusBadRequest, "Error generating token", err)
		log.Printf("Error generating token: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func handleError(c *gin.Context, statusCode int, logMessage string, err error) {
	c.JSON(statusCode, gin.H{"error": err.Error()})
	log.Printf("%s: %v", logMessage, err)
}
