package routes

import (
	"github.com/Tijanieneye10/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func homePage(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	//context.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {

	//get token
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized user"})
		return
	}

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})

}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "bad request"})
		return
	}

	event, err := models.GetSingleEvent(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Event not found"})
		return
	}

	context.JSON(http.StatusOK, event)

}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	}

	event, err := models.GetSingleEvent(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Event not found in database"})
		return
	}

	err = event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Event not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})

}

func updateEvent(context *gin.Context) {
	eventId, _ := strconv.ParseInt(context.Param("id"), 10, 64)

	_, err := models.GetSingleEvent(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	var event models.Event

	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		return
	}

	event.ID = int(eventId)

	err = event.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		return
	}

	context.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong", "event": event})

}
