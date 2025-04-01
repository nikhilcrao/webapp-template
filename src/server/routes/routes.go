package routes

import (
	"webapp/server/handlers"
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
	// protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/profile", handlers.GetProfile)

		handlers.RegisterCRUDHandlers[models.Account](
			handlers.HandlerConfig{
				BasePath:    "/accounts",
				RouterGroup: protected,
				UserScoped:  true,
			},
		)

		handlers.RegisterCRUDHandlers[models.Category](
			handlers.HandlerConfig{
				BasePath:    "/categories",
				RouterGroup: protected,
				UserScoped:  true,
			},
		)

		handlers.RegisterCRUDHandlers[models.Merchant](
			handlers.HandlerConfig{
				BasePath:    "/merchants",
				RouterGroup: protected,
				UserScoped:  true,
			},
		)

		handlers.RegisterCRUDHandlers[models.Transaction](
			handlers.HandlerConfig{
				BasePath:    "/transactions",
				RouterGroup: protected,
				UserScoped:  true,
			},
		)

		handlers.RegisterCRUDHandlers[models.Rule](
			handlers.HandlerConfig{
				BasePath:    "/rules",
				RouterGroup: protected,
				UserScoped:  true,
			},
		)
	}

	admin := api.Group("/admin")
	// Disabled for testing
	// admin.Use(middlewares.AuthMiddleware())
	{
		handlers.RegisterCRUDHandlers[models.AppUser](
			handlers.HandlerConfig{
				BasePath:    "/users",
				RouterGroup: admin,
				UserScoped:  false,
			},
		)
	}
}
