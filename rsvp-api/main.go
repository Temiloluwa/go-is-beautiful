package main

import (
	"github.com/gin-gonic/gin"
	"github.com/temiloluwa/rsvp-api/db"
	"github.com/temiloluwa/rsvp-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
