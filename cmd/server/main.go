package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/hrishabpachange/golang-mvc-boilerplate/config"
	"github.com/hrishabpachange/golang-mvc-boilerplate/database"
	"github.com/hrishabpachange/golang-mvc-boilerplate/routes"
)

func main() {
	//Load environment variables
	err := config.LoadEnv()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	//Initialize database connection
	database.Connect()

	//Initialize Gin router
	router := gin.Default()

	//Setup routes
	routes.SetupRoutes(router)

	//Start the server
	port := config.GetEnv("PORT", "8080")
	log.Printf("Server is running on port %s", port)

	err = router.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}