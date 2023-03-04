package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/otisnado/fofo-server/models"
)

func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func FindUser(c *gin.Context) {
	var user models.User
	user_id, _ := strconv.Atoi(c.Param("id"))

	if err := models.DB.Where("id = ?", user_id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func CreateUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: input.Name, Lastname: input.Lastname, Username: input.Username, Mail: input.Mail, Password: input.Password, Group: input.Group, State: input.State, CreatedAt: time.Now(), UpdatedAt: input.UpdatedAt}
	models.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	user_id, _ := strconv.Atoi(c.Param("id"))

	if err := models.DB.Where("id = ?", user_id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	user_id, _ := strconv.Atoi(c.Param("id"))

	if err := models.DB.Where("id = ?", user_id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
