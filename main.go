package main

import (
	"campaigns-restapi/auth"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() //* Initialize router with gin
	router.GET("/users", auth.GetUsers) //* Define router endpoint with HTTP request method and handler/controller
	router.Run() //* Running initialized router
}
