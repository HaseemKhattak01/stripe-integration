package routes

import (
	"github.com/HaseemKhattak01/stripe-integration/controllers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter(authService controllers.AuthService, validationService controllers.ValidationService) *Router {
	engine := gin.Default()

	engine.POST("/generate-token", func(c *gin.Context) {
		controllers.GenerateToken(c, authService, validationService)
	})

	return &Router{Engine: engine}
}
