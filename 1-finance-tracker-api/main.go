package main

import (
	"github.com/gin-gonic/gin"
	"github.com/temiloluwa/1-finance-tracker-api/db"
	"github.com/temiloluwa/1-finance-tracker-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
