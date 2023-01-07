package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var Db *gorm.DB

func Connection() *gorm.DB {
	godotenv.Load()
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v \n", err)
		return nil
	}
	fmt.Println("<==================Connection database successfully====================>", dsn)

	//db.AutoMigrate(&models.Category{})

	return db
}
func InitDb() *gorm.DB {
	return Connection()
}
