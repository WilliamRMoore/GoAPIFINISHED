package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventid, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID."})
		return
	}

	event, err := models.GetEventByID(eventid)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event!"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "IDK MAN"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registered toi event!"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventid, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event
	event.ID = eventid

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not unregister from event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Unregistered from event!"})
}
