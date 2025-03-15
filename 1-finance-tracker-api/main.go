package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temiloluwa/1-finance-tracker-api/db"
	"github.com/temiloluwa/1-finance-tracker-api/models"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

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

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.BindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}
	err = event.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}
