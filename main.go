package main

import (
	"campaigns-restapi/auth"
	"campaigns-restapi/campaign"
	"log"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	campaignRepository := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepository)
	campaignHandler := campaign.NewCampaignHandler(campaignService)

	router := gin.Default()
	firstVerAPI := router.Group("/api/v1")

	firstVerAPI.POST("/auth/signup", userHandler.SignUpHandler)
	firstVerAPI.POST("/auth/login", userHandler.LogInHandler)
	
	firstVerAPI.GET("/campaigns", campaignHandler.GetCampaigns)

	router.Run()
}