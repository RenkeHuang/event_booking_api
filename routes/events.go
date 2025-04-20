package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"event_booking/models"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch events."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": events})
}

func getEvent(context *gin.Context) {
	// Get the id from the URL
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch event by ID."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": event})
}

func createEvents(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		errText := fmt.Sprintf("Failed to parse request data: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": errText})
		return
	}

	userId := context.GetInt64("userId")
	event.UserID = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the event."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created."})
}

func updateEvent(context *gin.Context) {
	// Get the id from the URL
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch event by ID."})
		return
	}
	// Check if the event belongs to the user
	userId := context.GetInt64("userId")
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "No permission to update this event."})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		errText := fmt.Sprintf("Failed to parse request data: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": errText})
		return
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated."})
}

func deleteEvent(context *gin.Context) {
	// Get the id from the URL
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch event by ID."})
		return
	}
	// Check if the event belongs to the user
	userId := context.GetInt64("userId")
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "No permission to delete this event."})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted."})
}
