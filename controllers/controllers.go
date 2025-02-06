package controllers

import (
	"log"
	"net/http"

	"github.com/HaseemKhattak01/stripe-integration/models"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	GenerateToken(models.CardDetails) (string, error)
}

type ValidationService interface {
	ValidateCardDetails(models.CardDetails) error
}

func GenerateToken(c *gin.Context, authService AuthService, validationService ValidationService) {
	var cardDetails models.CardDetails
	if err := bindAndValidateCardDetails(c, &cardDetails, validationService); err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	token, err := authService.GenerateToken(cardDetails)
	if err != nil {
		handleError(c, http.StatusBadRequest, "Error generating token: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func bindAndValidateCardDetails(c *gin.Context, cardDetails *models.CardDetails, validationService ValidationService) error {
	if err := c.ShouldBindJSON(cardDetails); err != nil {
		return "Error binding JSON: " + err.Error()
	}

	if err := validationService.ValidateCardDetails(*cardDetails); err != nil {
		return "Validation error: " + err.Error()
	}

	return nil
}

func handleError(c *gin.Context, statusCode int, logMessage string) {
	c.JSON(statusCode, gin.H{"error": logMessage})
	log.Println(logMessage)
}
