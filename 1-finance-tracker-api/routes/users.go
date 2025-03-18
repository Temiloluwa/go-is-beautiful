package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/temiloluwa/1-finance-tracker-api/models"
	"github.com/temiloluwa/1-finance-tracker-api/utils"
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

func login(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful.", "token": token})
}
