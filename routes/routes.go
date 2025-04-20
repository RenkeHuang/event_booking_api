package routes

import (
	"event_booking/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Setup endpoints: GET, POST, PUT, DELETE, PATCH
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) // e.g. /events/1, /events/2

	// Auth required
	authenticated := server.Group("/")
	// Auth middleware will always be executed before all handlers in this group
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvents)
	authenticated.PUT("/events/:id", updateEvent)                    // e.g. /events/1, /events/2
	authenticated.DELETE("/events/:id", deleteEvent)                 // e.g. /events/1, /events/2
	authenticated.POST("/events/:id/register", registerForEvent)     // e.g. /events/1/register, /events/2/register
	authenticated.DELETE("/events/:id/register", cancelRegistration) // e.g. /events/1/register, /events/2/register
	authenticated.GET("/events/:id/register", getAllRegistrationsForEvent)

	// user
	server.POST("/signup", signUp)
	server.POST("/login", login)
}
