package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/otisnado/fofo-server/models"
)

// FindProjects		godoc
// @Summary			FindProjects
// @Schemes
// @Description		Bulk all fofo-server's projects
// @Tags			Projects
// @Produce			json
// @Param			Authorization		header	string	true	"JWT without bearer"
// @Success			200		{object}	models.SuccessFindProjects
// @Failure			401		{object}	models.ErrorMessage
// @Router			/projects	[get]
func FindGroups(c *gin.Context) {
	var groups []models.Group
	models.DB.Find(&groups)

	c.JSON(http.StatusOK, gin.H{"data": groups})
}

// FindProject	godoc
// @Summary		FindProject
// @Schemes
// @Description	Find a project with given id in path
// @Tags		Projects
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		id		path		int		true	"Project ID"
// @Success		200		{object}	models.SuccessFindProject
// @Failure		400,401	{object}	models.ErrorMessage
// @Router		/projects/{id}			[get]
func FindGroup(c *gin.Context) {
	var group models.Group
	group_id, _ := strconv.Atoi(c.Param("id"))

	if err := models.DB.Where("id = ?", group_id).First(&group).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Group not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": group})
}

// CreateProjects	godoc
// @Summary			CreateProject
// @Schemes
// @Description		Create project with models.Project model
// @Tags			Projects
// @Produce			json
// @Param			Authorization		header	string	true	"JWT without bearer"
// @Param			project	body		models.Project	true	"Project data"
// @Success			200		{object}	models.SuccessProjectCreation
// @Failure			400,401	{object}	models.ErrorMessage
// @Router			/projects	[post]
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

// UpdateProject	godoc
// @Summary			UpdateProject
// @Schemes
// @Description		Update project with models.Project model
// @Tags			Projects
// @Produce			json
// @Param			Authorization		header	string	true	"JWT without bearer"
// @Param			project			body		models.Project			true	"Project data"
// @Success			200		{object}	models.SuccessProjectUpdate
// @Failure			400,401	{object}	models.ErrorMessage
// @Router			/projects	[patch]
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

// DeleteProject	godoc
// @Summary			DeleteProject
// @Schemes
// @Description		Delete project using id
// @Tags			Projects
// @Produce			json
// @Param			Authorization		header	string	true	"JWT without bearer"
// @Param			id		path		int				true	"Project ID"
// @Success			200		{object}	models.SuccessProjectDelete
// @Failure			400,401	{object}	models.ErrorMessage
// @Router			/projects/{id}	[delete]
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
