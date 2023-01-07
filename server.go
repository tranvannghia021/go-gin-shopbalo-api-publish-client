package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-api-publish/routers"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	environment := os.Getenv("ENV")

	fmt.Println("<======>ENVIRONMENT=", environment)

	server := gin.New()
	server.Use()
	apiGroup := server.Group("/api")
	routers.ApiRouter(apiGroup)
	server.Run(":8000")
}
