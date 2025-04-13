package main

import (
	"event_booking/db"
	"event_booking/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB() // Initialize the database
	server := gin.Default()

	routes.RegisterRoutes(server) // Register the routes

	server.Run(":8080") //localhost:8080
}
