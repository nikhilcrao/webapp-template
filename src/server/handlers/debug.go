package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleHealthz(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":   "ok",
		"timestamp": time.Now(),
	})
}

func HandleDebugz(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":   "ok",
		"timestamp": time.Now(),
	})
}
