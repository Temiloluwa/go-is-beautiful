package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/temiloluwa/1-finance-tracker-api/models"
)

func signup(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data."})
		return
	}

	err = user.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully."})

}
