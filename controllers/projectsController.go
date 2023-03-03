package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/otisnado/fofo-server/models"
)

func FindProjects(c *gin.Context) {
	var projects []models.Project
	models.DB.Find(&projects)

	c.JSON(http.StatusOK, gin.H{"data": projects})
}

func FindProject(c *gin.Context) {
	var project models.Project
	project_id, _ := strconv.Atoi(c.Param("id"))

	if err := models.DB.Where("id = ?", project_id).First(&project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": project})
}

func CreateProject(c *gin.Context) {
	var input models.InputProject
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project := models.Project{Name: input.Name, Created_by: input.Created_by, Language: input.Language}
	models.DB.Create(&project)
	c.JSON(http.StatusCreated, gin.H{"data": project})
}

func UpdateProject(c *gin.Context) {
	var project models.Project
	project_id, _ := strconv.Atoi(c.Param("id"))

	if err := models.DB.Where("id = ?", project_id).First(&project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project not found!"})
		return
	}

	var input models.UpdateProjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&project).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": project})
}

func DeleteProject(c *gin.Context) {
	var project models.Project
	project_id, _ := strconv.Atoi(c.Param("id"))

	if err := models.DB.Where("id = ?", project_id).First(&project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project not found!"})
		return
	}

	models.DB.Delete(&project)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
