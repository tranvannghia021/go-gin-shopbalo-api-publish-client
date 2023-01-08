package main

import (
	"fmt"
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
	helpers.LoggerFile()
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	environment := os.Getenv("ENV")

	fmt.Println("<======>ENVIRONMENT=", environment)

	server := gin.Default()
	server.Use(middlewares.CorsMiddleware())
	apiGroup := server.Group("/api/client")
	routers.ApiRouter(apiGroup)
	server.Run(":8001")
}
