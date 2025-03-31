package routes

import (
	"webapp/server/handlers"
	"webapp/server/middlewares"
	"webapp/server/models"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	api.GET("/health", handlers.HandleHealth)

	auth := api.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)

		auth.GET("/google", handlers.GoogleLogin)
		auth.GET("/google/callback", handlers.GoogleCallback)
	}

	protected := api.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/profile", handlers.GetProfile)
	}

	admin := api.Group("/admin")
	// Disabled for testing
	// admin.Use(middlewares.AuthMiddleware())
	{
		handlers.RegisterCRUDHandlers[models.User](
			handlers.HandlerConfig{
				BasePath:    "/users",
				RouterGroup: admin,
			},
		)
	}
}
