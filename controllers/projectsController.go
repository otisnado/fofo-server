package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/otisnado/fofo-server/models"
)

// FindGroups		godoc
// @Summary			FindGroups
// @Schemes
// @Description		Bulk all fofo-server's groups
// @Tags			Groups
// @Produce			json
// @Param			Authorization		header	string	true	"JWT without bearer"
// @Success			200		{object}	models.SuccessFindGroups
// @Failure			401		{object}	models.ErrorMessage
// @Router			/groups	[get]
func FindProjects(c *gin.Context) {
	var projects []models.Project
	models.DB.Find(&projects)

	c.JSON(http.StatusOK, gin.H{"data": projects})
}

// FindGroup	godoc
// @Summary		FindGroup
// @Schemes
// @Description	Find a group with given id in path
// @Tags		Groups
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		id		path		int		true	"Group ID"
// @Success		200		{object}	models.SuccessFindGroup
// @Failure		400,401	{object}	models.ErrorMessage
// @Router		/groups/{id}			[get]
func FindProject(c *gin.Context) {
	var project models.Project
	project_id, _ := strconv.Atoi(c.Param("id"))

	if err := models.DB.Where("id = ?", project_id).First(&project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": project})
}

// CreateGroup	godoc
// @Summary		CreateGroup
// @Schemes
// @Description	Create group with models.Group model
// @Tags		Groups
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		group	body		models.Group		true	"Group data"
// @Success		200		{object}	models.SuccessGroupCreation
// @Failure		400,401	{object}	models.ErrorMessage
// @Router		/groups	[post]
func CreateProject(c *gin.Context) {
	var input models.Project
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project := models.Project{Name: input.Name, Created_by: input.Created_by, Language: input.Language, CreatedAt: time.Now(), UpdatedAt: input.UpdatedAt}
	models.DB.Create(&project)
	c.JSON(http.StatusCreated, gin.H{"data": project})
}

// UpdateGroup	godoc
// @Summary		UpdateGroup
// @Schemes
// @Description	Update group with models.Group model
// @Tags		Groups
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		group	body		models.Group			true	"Group data"
// @Success		200		{object}	models.SuccessGroupUpdate
// @Failure		400,401	{object}	models.ErrorMessage
// @Router		/groups/{id}	[patch]
func UpdateProject(c *gin.Context) {
	var project models.Project
	project_id, _ := strconv.Atoi(c.Param("id"))

	if err := models.DB.Where("id = ?", project_id).First(&project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project not found!"})
		return
	}

	var input models.Project
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&project).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": project})
}

// DeleteGroup	godoc
// @Summary		DeleteGroup
// @Schemes
// @Description	Delete group with using id
// @Tags		Groups
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		id		path		int				true	"Group ID"
// @Success		200		{object}	models.SuccessGroupDelete
// @Failure		400,401	{object}	models.ErrorMessage
// @Router		/groups/{id}	[delete]
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
