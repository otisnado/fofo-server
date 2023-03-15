package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/otisnado/nepackage/models"
	"github.com/otisnado/nepackage/services"
)

// FindProjects		godoc
// @Summary			FindProjects
// @Schemes
// @Description		Bulk all nepackage's projects
// @Tags			Projects
// @Produce			json
// @Param			Authorization		header	string	true	"JWT without bearer"
// @Success			200		{object}	models.SuccessFindProjects
// @Failure			401,500	{object}	models.ErrorMessage
// @Router			/projects	[get]
func FindGroups(c *gin.Context) {
	groups, err := services.GetGroups()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": groups})
}

// FindProject	godoc
// @Summary		FindProject
// @Schemes
// @Description	Find a project with given id in path
// @Tags		Projects
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		id			path		int		true	"Project ID"
// @Success		200			{object}	models.SuccessFindProject
// @Failure		400,401,404	{object}	models.ErrorMessage
// @Router		/projects/{id}			[get]
func FindGroup(c *gin.Context) {
	group_id, _ := strconv.Atoi(c.Param("id"))

	group, err := services.GetGroupById(uint(group_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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
// @Param			project		body		models.Project	true	"Project data"
// @Success			200			{object}	models.SuccessProjectCreation
// @Failure			400,401,500	{object}	models.ErrorMessage
// @Router			/projects	[post]
func CreateGroup(c *gin.Context) {
	var input models.Group
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := services.CreateGroup(&input)
	if !created {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": &input})
}

// UpdateProject	godoc
// @Summary			UpdateProject
// @Schemes
// @Description		Update project with models.Project model
// @Tags			Projects
// @Produce			json
// @Param			Authorization		header	string	true	"JWT without bearer"
// @Param			project			body		models.Project			true	"Project data"
// @Success			200				{object}	models.SuccessProjectUpdate
// @Failure			400,401,404,500	{object}	models.ErrorMessage
// @Router			/projects	[patch]
func UpdateGroup(c *gin.Context) {
	group_id, _ := strconv.Atoi(c.Param("id"))

	var input models.GroupUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := services.GetGroupById(uint(group_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	groupUpdated, err := services.UpdateGroup(uint(group_id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": groupUpdated})
}

// DeleteProject	godoc
// @Summary			DeleteProject
// @Schemes
// @Description		Delete project using id
// @Tags			Projects
// @Produce			json
// @Param			Authorization		header	string	true	"JWT without bearer"
// @Param			id			path		int				true	"Project ID"
// @Success			200			{object}	models.SuccessProjectDelete
// @Failure			401,404,500	{object}	models.ErrorMessage
// @Router			/projects/{id}	[delete]
func DeleteGroup(c *gin.Context) {
	group_id, _ := strconv.Atoi(c.Param("id"))

	_, err := services.GetGroupById(uint(group_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	state, err := services.DeleteGroup(uint(group_id))
	if !state {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": state})
}
