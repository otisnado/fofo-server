package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/otisnado/nepackage/models"
	"github.com/otisnado/nepackage/services"
)

// FindGroups		godoc
// @Summary			FindGroups
// @Schemes
// @Description		Bulk all nepackage's groups
// @Tags			Groups
// @Produce			json
// @Param			Authorization		header	string	true	"JWT without bearer"
// @Success			200		{object}	models.SuccessFindGroups
// @Failure			401,500	{object}	models.ErrorMessage
// @Router			/groups	[get]
func FindProjects(c *gin.Context) {
	projects, err := services.GetProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": projects})
}

// FindGroup	godoc
// @Summary		FindGroup
// @Schemes
// @Description	Find a group with given id in path
// @Tags		Groups
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		id			path		int		true	"Group ID"
// @Success		200			{object}	models.SuccessFindGroup
// @Failure		400,401,404	{object}	models.ErrorMessage
// @Router		/groups/{id}			[get]
func FindProject(c *gin.Context) {
	project_id, _ := strconv.Atoi(c.Param("id"))

	project, err := services.GetProjectById(uint(project_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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
// @Param		group		body		models.Group		true	"Group data"
// @Success		200			{object}	models.SuccessGroupCreation
// @Failure		400,401,500	{object}	models.ErrorMessage
// @Router		/groups	[post]
func CreateProject(c *gin.Context) {
	var input models.Project
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := services.CreateProject(&input)
	if !created {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": &input})
}

// UpdateGroup	godoc
// @Summary		UpdateGroup
// @Schemes
// @Description	Update group with models.Group model
// @Tags		Groups
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		group			body		models.Group			true	"Group data"
// @Success		200				{object}	models.SuccessGroupUpdate
// @Failure		400,401,404,500	{object}	models.ErrorMessage
// @Router		/groups/{id}	[patch]
func UpdateProject(c *gin.Context) {
	project_id, _ := strconv.Atoi(c.Param("id"))

	var input models.ProjectUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := services.GetProjectById(uint(project_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	projectUpdated, err := services.UpdateProject(uint(project_id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": projectUpdated})
}

// DeleteGroup	godoc
// @Summary		DeleteGroup
// @Schemes
// @Description	Delete group with using id
// @Tags		Groups
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		id			path		int				true	"Group ID"
// @Success		200			{object}	models.SuccessGroupDelete
// @Failure		401,404,500	{object}	models.ErrorMessage
// @Router		/groups/{id}	[delete]
func DeleteProject(c *gin.Context) {
	project_id, _ := strconv.Atoi(c.Param("id"))

	_, err := services.GetProjectById(uint(project_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	state, err := services.DeleteProject(uint(project_id))
	if !state {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": state})
}
