package routes

import (
	"fmt"
	"net/http"

	"event_booking/models"
	"event_booking/utils"

	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		errText := fmt.Sprintf("Failed to parse request data: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": errText})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user: " + err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created."})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		errText := fmt.Sprintf("Failed to parse request data: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": errText})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to authenticate user: " + err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token: " + err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "User login successful.",
		"token":   token,
	})
}
