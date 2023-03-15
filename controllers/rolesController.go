package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/otisnado/nepackage/models"
	"github.com/otisnado/nepackage/services"
)

// FindRoles		godoc
// @Summary			FindRoles
// @Schemes
// @Description		Bulk all nepackage's roles
// @Tags			Roles
// @Produce			json
// @Param			Authorization		header	string	true	"JWT without bearer"
// @Success			200		{object}	models.SuccessFindRoles
// @Failure			401,500	{object}	models.ErrorMessage
// @Router			/roles	[get]
func FindRoles(c *gin.Context) {
	roles, err := services.GetRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": roles})
}

// FindRole	godoc
// @Summary		FindRole
// @Schemes
// @Description	Find a role with given id in path
// @Tags		Roles
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		id			path		int		true	"Role ID"
// @Success		200			{object}	models.SuccessFindRole
// @Failure		400,401,404	{object}	models.ErrorMessage
// @Router		/roles/{id}			[get]
func FindRole(c *gin.Context) {
	project_id, _ := strconv.Atoi(c.Param("id"))

	project, err := services.GetRoleById(uint(project_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": project})
}

// CreateRole	godoc
// @Summary		CreateRole
// @Schemes
// @Description	Create role with models.Role model
// @Tags		Roles
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		role		body		models.Role		true	"Role data"
// @Success		200			{object}	models.SuccessRoleCreation
// @Failure		400,401,500	{object}	models.ErrorMessage
// @Router		/roles	[post]
func CreateRole(c *gin.Context) {
	var input models.Role
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := services.CreateRole(&input)
	if !created {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": &input})
}

// UpdateRole	godoc
// @Summary		UpdateRole
// @Schemes
// @Description	Update role with models.Role model
// @Tags		Roles
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		role			body		models.Role			true	"Role data"
// @Success		200				{object}	models.SuccessRoleUpdate
// @Failure		400,401,404,500	{object}	models.ErrorMessage
// @Router		/roles/{id}	[patch]
func UpdateRole(c *gin.Context) {
	project_id, _ := strconv.Atoi(c.Param("id"))

	var input models.RoleUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := services.GetRoleById(uint(project_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	projectUpdated, err := services.UpdateRole(uint(project_id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": projectUpdated})
}

// DeleteRole	godoc
// @Summary		DeleteRole
// @Schemes
// @Description	Delete role with using id
// @Tags		Roles
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		id			path		int				true	"Role ID"
// @Success		200			{object}	models.SuccessRoleDelete
// @Failure		401,404,500	{object}	models.ErrorMessage
// @Router		/roles/{id}	[delete]
func DeleteRole(c *gin.Context) {
	project_id, _ := strconv.Atoi(c.Param("id"))

	_, err := services.GetRoleById(uint(project_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	state, err := services.DeleteRole(uint(project_id))
	if !state {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": state})
}

// FindRolesPolicies	godoc
// @Summary				FindRolesPolicies
// @Schemes
// @Description	Find all policies associated to given role with role id in path
// @Tags		Roles
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		id			path		int		true	"Role ID"
// @Success		200			{object}	models.SuccessFindPolicies
// @Failure		400,401,404	{object}	models.ErrorMessage
// @Router		/roles/{id}			[get]
func FindRolesPolicies(c *gin.Context) {
	role_id, _ := strconv.Atoi(c.Param("id"))

	policies, err := services.GetPoliciesByRoleId(uint(role_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": policies})
}
