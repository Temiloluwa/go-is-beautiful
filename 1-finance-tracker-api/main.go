package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/temiloluwa/1-finance-tracker-api/models"
)

func main() {
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func getEvents(c *gin.Context) {
	events := models.GetAllEvents()

	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.BindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}
	event.ID = 1
	event.UserId = 1
	event.Save()
	c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}
