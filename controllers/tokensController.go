package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otisnado/fofo-server/auth"
	"github.com/otisnado/fofo-server/models"
)

func GenerateToken(c *gin.Context) {
	var request models.TokenRequest
	var user models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// Check if mail exists and password is correct
	record := models.DB.Where("mail = ?", request.Mail).First(&user)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		c.Abort()
		return
	}

	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		c.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(user.Mail, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
