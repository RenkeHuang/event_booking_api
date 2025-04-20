package routes

import (
	"event_booking/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
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

	isRegistered, _ := event.IsUserRegistered(userId)
	if isRegistered {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Already registered for this event."})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register for event."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Registered for event."})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id."})
		return
	}

	var event models.Event
	event.ID = eventId
	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel registration."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Cancelled registration."})
}

func getAllRegistrationsForEvent(context *gin.Context) {
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

	// Only allow the user who created the event to see the registrations
	userId := context.GetInt64("userId")
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "No permission to see registrations for this event."})
		return
	}

	userIds, err := event.GetRegistrations()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch registrations."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"Registered UserIds": userIds})
}
