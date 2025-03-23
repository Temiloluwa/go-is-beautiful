package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temiloluwa/rsvp-api/models"
)

func getEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}
	c.JSON(http.StatusOK, event)
}

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
		return
	}

	c.JSON(http.StatusOK, events)
}

// protected route
func createEvent(c *gin.Context) {

	var event models.Event
	err := c.BindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	event.UserId = c.GetInt("userId")
	err = event.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

func updateEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return

	}

	userId := c.GetInt("userId")
	event, err := models.GetEventByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	if userId != event.UserId {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	err = c.BindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	event.ID = id
	err = event.Update()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}
	userId := c.GetInt("userId")
	event, err := models.GetEventByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	// user that created event must be the one to update it
	if userId != event.UserId {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	err = event.Delete()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
