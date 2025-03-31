package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
	"webapp/server"
	"webapp/server/config"
	"webapp/server/database"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/joho/godotenv"
)

func main() {
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		glog.Warning(".env not found")
	}

	cfg := config.LoadConfig()

	err = database.Init()
	if err != nil {
		glog.Fatal(err)
	}

	router := gin.New()
	router.Use(gin.Logger())
	// router.Use(gin.Recovery())
	router.SetTrustedProxies(nil)

	corsConfig := cors.DefaultConfig()
	// corsConfig.AllowedOrigins = []string{"*"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowedHeaders("Authorization")
	router.Use(cors.New(corsConfig))

	router.Static("/assets", "./client/dist/assets")
	router.Static("/favicon.ico", "./client/dist/favicon.ico")
	router.Static("/manifest.json", "./client/dist/manifest.json")

	router.NoRoute(func(ctx *gin.Context) {
		if strings.HasPrefix(ctx.Request.URL.Path, "/api") {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "API endpoint not found"})
			return
		}
		// Serve index.html for client-side routing
		ctx.File("./client/dist/index.html")
	})

	server, err := server.NewServer(router)
	if err != nil {
		glog.Fatal(err)
	}

	server.Run(fmt.Sprintf("%s:%s", cfg.Addr, cfg.Port))
}
