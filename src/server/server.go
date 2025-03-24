package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer(r *gin.Engine) (*Server, error) {
	s := Server{router: r}
	s.setupRoutes()
	return &s, nil
}

func (s *Server) setupRoutes() {
	s.router.GET("/ping", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "pong"}) })
}

func (s *Server) Run(addr string) {
	s.router.Run(addr)
}
