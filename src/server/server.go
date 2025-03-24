package server

import (
	"webapp/server/routes"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer(r *gin.Engine) (*Server, error) {
	s := Server{router: r}
	routes.RegisterRoutes(s.router)
	return &s, nil
}

func (s *Server) Run(addr string) {
	s.router.Run(addr)
}
