package main

import (
	"campaigns-restapi/auth"
	"campaigns-restapi/campaign"
	"campaigns-restapi/middleware"
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

	router.Static("/campaign-images", "./campaign-images") //* Configure campaign-images accessible via request endpoint
	router.Static("/images", "./images")

	firstVerAPI := router.Group("/api/v1")

	firstVerAPI.POST("/auth/signup", userHandler.SignUpHandler)
	firstVerAPI.POST("/auth/login", userHandler.LogInHandler)
	
	firstVerAPI.GET("/campaigns", campaignHandler.GetCampaigns)
	firstVerAPI.GET("/campaigns/:id", campaignHandler.GetSpecifiedCampaign) //* Configure endpoint with params 'id'
	firstVerAPI.POST("/campaigns", middleware.AuthorizationMiddleware(userService), campaignHandler.CreateNewCampaign)

	router.Run()
}