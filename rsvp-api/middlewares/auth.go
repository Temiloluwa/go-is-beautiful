package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/temiloluwa/rsvp-api/utils"
)

func Authenticate(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return

	}

	c.Set("userId", userId)
	c.Next()

}
