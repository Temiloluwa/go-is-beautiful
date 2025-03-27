package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/temiloluwa/rsvp-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	// Register all routes here

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.GET("/users", getUsers)
	server.POST("/signup", signup)
	server.POST("/login", login)

}
