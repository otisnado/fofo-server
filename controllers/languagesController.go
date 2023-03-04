package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/otisnado/fofo-server/models"
)

func FindLanguages(c *gin.Context) {
	var languages []models.Language
	models.DB.Find(&languages)

	c.JSON(http.StatusOK, gin.H{"data": languages})
}

func FindLanguage(c *gin.Context) {
	var language models.Language
	language_id, _ := strconv.Atoi(c.Param("id"))

	if err := models.DB.Where("id = ?", language_id).First(&language).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Language not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": language})
}

func CreateLanguage(c *gin.Context) {
	var input models.Language
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project := models.Language{Name: input.Name, Created_by: input.Created_by, CreatedAt: time.Now(), UpdatedAt: input.UpdatedAt}
	models.DB.Create(&project)
	c.JSON(http.StatusCreated, gin.H{"data": project})
}

func UpdateLanguage(c *gin.Context) {
	var language models.Language
	project_id, _ := strconv.Atoi(c.Param("id"))

	if err := models.DB.Where("id = ?", project_id).First(&language).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Language not found!"})
		return
	}

	var input models.Language
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&language).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": language})
}
