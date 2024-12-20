package main

import (
	"log"

	"coditas-task/handlers"
	"coditas-task/middleware"
	"coditas-task/validators" // Import validators package

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	validate := validator.New()
	validators.RegisterValidators(validate)

	// Initialize Gin router
	r := gin.Default()

	// Apply middleware for logging latency
	r.Use(middleware.LogLatency)

	// Initialize the handler with the shared validator
	userHandler := handlers.NewUserHandler(validate)

	// Define the POST route for creating a user
	r.POST("/user", userHandler.CreateUser)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
