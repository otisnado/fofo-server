package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/otisnado/fofo-server/models"
)

func FindGroups(c *gin.Context) {
	var groups []models.Group
	models.DB.Find(&groups)

	c.JSON(http.StatusOK, gin.H{"data": groups})
}

func FindGroup(c *gin.Context) {
	var group models.Group
	group_id, _ := strconv.Atoi(c.Param("id"))

	if err := models.DB.Where("id = ?", group_id).First(&group).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Group not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": group})
}

func CreateGroup(c *gin.Context) {
	var input models.Group
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group := models.Group{Name: input.Name, CreatedAt: time.Now(), UpdatedAt: input.UpdatedAt}
	models.DB.Create(&group)
	c.JSON(http.StatusCreated, gin.H{"data": group})
}

func UpdateGroup(c *gin.Context) {
	var group models.Group
	group_id, _ := strconv.Atoi(c.Param("id"))

	if err := models.DB.Where("id = ?", group_id).First(&group).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Group not found!"})
		return
	}

	var input models.Group
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&group).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": group})
}

func DeleteGroup(c *gin.Context) {
	var group models.Group
	group_id, _ := strconv.Atoi(c.Param("id"))

	if err := models.DB.Where("id = ?", group_id).First(&group).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Group not found!"})
		return
	}

	models.DB.Delete(&group)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
