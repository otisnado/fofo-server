package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/otisnado/fofo-server/models"
	"github.com/otisnado/fofo-server/utils"
)

// FindUsers		godoc
// @Summary			FindUsers
// @Schemes
// @Description		Bulk all fofo-server's users
// @Tags			Users
// @Produce			json
// @Param			Authorization		header	string	true	"JWT without bearer"
// @Success			200		{object}	models.SuccessFindUsers
// @Failure			401		{object}	models.ErrorMessage
// @Router			/users	[get]
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// FindUser		godoc
// @Summary		FindUser
// @Schemes
// @Description	Find a user with given id in path
// @Tags		Users
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		id		path		int		true	"User ID"
// @Success		200		{object}	models.SuccessFindUser
// @Failure		400,401	{object}	models.ErrorMessage
// @Router		/users/{id}			[get]
func FindUser(c *gin.Context) {
	var user models.User
	user_id, _ := strconv.Atoi(c.Param("id"))

	if err := models.DB.Where("id = ?", user_id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// CreateUser	godoc
// @Summary		CreateUser
// @Schemes
// @Description	Create user with models.User model
// @Tags		Users
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		user	body		models.User		true	"User data"
// @Success		200		{object}	models.SuccessUserCreation
// @Failure		400,401	{object}	models.ErrorMessage
// @Router		/users	[post]
func CreateUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.CheckIfMailExists(input.Mail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.CheckIfUsernameExists(input.Username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := input.HashPassword(input.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	user := models.User{Name: input.Name, Lastname: input.Lastname, Username: input.Username, Mail: input.Mail, Password: input.Password, Group: input.Group, State: input.State, CreatedAt: time.Now(), UpdatedAt: input.UpdatedAt}
	models.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"data": user})
}

// UpdateUser	godoc
// @Summary		UpdateUser
// @Schemes
// @Description	Update user with models.User model
// @Tags		Users
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		user	body		models.User			true	"User data"
// @Success		200		{object}	models.SuccessUserUpdate
// @Failure		400,401	{object}	models.ErrorMessage
// @Router		/users/{id}	[patch]
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

// DeleteUser	godoc
// @Summary		DeleteUser
// @Schemes
// @Description	Delete user with using id
// @Tags		Users
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		id		path		int				true	"User ID"
// @Success		200		{object}	models.SuccessUserDelete
// @Failure		400,401	{object}	models.ErrorMessage
// @Router		/users/{id}	[delete]
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
