package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"campaigns-restapi/auth"
)

func main() {
	dsn := "root:rootPassword!@tcp(127.0.0.1:3306)/golang-bwa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if (err != nil) {
		log.Fatal(err.Error())
	}

	log.Println("Successfully connected to database!")

	userRepository := auth.NewRepository(db)
	userService := auth.NewService(userRepository)
	userHandler := auth.NewUserHandler(userService)

	router := gin.Default()
	firstVerAPI := router.Group("/api/v1")

	firstVerAPI.POST("/auth/signup", userHandler.SignUpHandler)
	firstVerAPI.POST("/auth/login", userHandler.LogInHandler)

	router.Run()
}