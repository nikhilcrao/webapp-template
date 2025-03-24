package routes

import (
	"webapp/server/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")

	api.GET("/debugz", handlers.HandleDebugz)
	api.GET("/healthz", handlers.HandleHealthz)

	// TODO: Auth
	// TODO: Protected
}
