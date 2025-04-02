package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleHealth(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message":   "ok",
		"timestamp": time.Now(),
	})
}

func HandleNotImplemented(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusNotImplemented, gin.H{})
}
