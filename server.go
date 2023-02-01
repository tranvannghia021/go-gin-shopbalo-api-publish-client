package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-api-publish/app/helpers"
	"go-api-publish/app/middlewares"
	_ "go-api-publish/app/middlewares"
	"go-api-publish/routers"
	"log"
	"os"
)

func main() {
	//config.InitDb()
	helpers.Logger()
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	modeApp := os.Getenv("APP_DEBUG")
	gin.SetMode(gin.DebugMode)
	if modeApp != "true" {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()
	server.SetTrustedProxies(nil)
	server.Use(middlewares.CorsMiddleware())
	apiGroup := server.Group("/api/client")
	routers.ApiRouter(apiGroup)
	server.Run(":8001")
}
