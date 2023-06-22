package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/sujeetregmi/gin-note-app/models"
)

func GetUserFromRequest(c *gin.Context) *models.User {
	userID := c.GetUint64("user_id")
	var currentUser *models.User
	if userID > 0 {
		currentUser = models.UserFind(userID)
	} else {
		currentUser = nil
	}
	return currentUser
}
