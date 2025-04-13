package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Setup endpoints: GET, POST, PUT, DELETE, PATCH
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) // e.g. /events/1, /events/2

	// Auth required
	server.POST("/events", createEvents)
	server.PUT("/events/:id", updateEvent)    // e.g. /events/1, /events/2
	server.DELETE("/events/:id", deleteEvent) // e.g. /events/1, /events/2

	// user
	server.POST("/signup", signUp)
	server.POST("/login", login)
}
