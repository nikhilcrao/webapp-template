package server

import "github.com/gin-gonic/gin"

type Server struct {
	router *gin.Engine
}

func NewServer(r *gin.Engine) (*Server, error) {
	server := Server{router: r}
	return &server, nil
}

func (s *Server) Run(addr string) {
	s.router.Run(addr)
}
