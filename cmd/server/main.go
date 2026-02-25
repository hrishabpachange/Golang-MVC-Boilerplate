// @title Golang MVC Boilerplate API
// @version 1.0
// @description Professional MVC Boilerplate using Gin
// @host localhost:8080
// @BasePath /

package main

import (
	"log"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/hrishabpachange/golang-mvc-boilerplate/docs"

	"github.com/hrishabpachange/golang-mvc-boilerplate/config"
	"github.com/hrishabpachange/golang-mvc-boilerplate/database"
	"github.com/hrishabpachange/golang-mvc-boilerplate/routes"
)

func main() {

	// Load environment variables
	err := config.LoadEnv()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	// Initialize database connection
	database.Connect()

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router)

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start the server
	port := config.GetEnv("PORT", "8080")
	log.Printf("Server is running on port %s", port)

	err = router.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}