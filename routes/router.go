package routes

import (
	"github.com/HaseemKhattak01/stripe-integration/controllers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter(authService controllers.AuthService, validationService controllers.ValidationService) *Router {
	router := &Router{Engine: gin.Default()}
	router.initializeRoutes(authService, validationService)
	return router
}

func (router *Router) initializeRoutes(authService controllers.AuthService, validationService controllers.ValidationService) {
	router.Engine.POST("/generate-token", func(c *gin.Context) {
		controllers.GenerateToken(c, authService, validationService)
	})
}
