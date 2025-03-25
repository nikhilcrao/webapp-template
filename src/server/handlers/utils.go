package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":   "ok",
		"timestamp": time.Now(),
	})
}

func HandleNotImplemented(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{})
}
