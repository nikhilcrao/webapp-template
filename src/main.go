package main

import (
	"webapp/server"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	s, err := server.NewServer(r)
	if err != nil {
		panic(err)
	}

	s.Run(":8080")
}
