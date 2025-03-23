package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temiloluwa/rsvp-api/models"
)

func registerForEvent(c *gin.Context) {
	userId := c.GetInt("userId")
	eventId, err := strconv.Atoi(c.Param("id"))
	event, err := models.GetEventByID(eventId)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	err = c.BindJSON(&event)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Registered for event"})
}

func cancelRegistration(c *gin.Context) {
	userId := c.GetInt("userId")
	eventId, err := strconv.Atoi(c.Param("id"))
	event, err := models.GetEventByID(eventId)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	err = c.BindJSON(&event)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = event.CancelRegistration(userId)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration for event"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Cancelled registration for event"})
}
