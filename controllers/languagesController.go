package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/otisnado/fofo-server/models"
)

// FindLanguages		godoc
// @Summary			FindLanguages
// @Schemes
// @Description		Bulk all fofo-server's languages
// @Tags			Languages
// @Produce			json
// @Param			Authorization		header	string	true	"JWT without bearer"
// @Success			200		{object}	models.SuccessFindLanguages
// @Failure			401		{object}	models.ErrorMessage
// @Router			/languages	[get]
func FindLanguages(c *gin.Context) {
	var languages []models.Language
	models.DB.Find(&languages)

	c.JSON(http.StatusOK, gin.H{"data": languages})
}

// FindLanguage		godoc
// @Summary		FindLanguage
// @Schemes
// @Description	Find a language with given id in path
// @Tags		Languages
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		id		path		int		true	"Language ID"
// @Success		200		{object}	models.SuccessFindLanguage
// @Failure		400,401	{object}	models.ErrorMessage
// @Router		/languages/{id}			[get]
func FindLanguage(c *gin.Context) {
	var language models.Language
	language_id, _ := strconv.Atoi(c.Param("id"))

	if err := models.DB.Where("id = ?", language_id).First(&language).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Language not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": language})
}

// CreateLanguage	godoc
// @Summary		CreateLanguage
// @Schemes
// @Description	Create language with models.Language model
// @Tags		Languages
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		language	body		models.Language		true	"Language data"
// @Success		200		{object}	models.SuccessLanguageCreation
// @Failure		400,401	{object}	models.ErrorMessage
// @Router		/languages	[post]
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

// UpdateLanguage	godoc
// @Summary		UpdateLanguage
// @Schemes
// @Description	Update language with models.Language model
// @Tags		Languages
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		language	body		models.Language			true	"Language data"
// @Success		200		{object}	models.SuccessLanguageUpdate
// @Failure		400,401	{object}	models.ErrorMessage
// @Router		/languages/{id}	[patch]
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

// DeleteLanguage	godoc
// @Summary		DeleteLanguage
// @Schemes
// @Description	Delete language with using id
// @Tags		Languages
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		id		path		int				true	"Language ID"
// @Success		200		{object}	models.SuccessLanguageDelete
// @Failure		400,401	{object}	models.ErrorMessage
// @Router		/languages/{id}	[delete]
func DeleteLanguage(c *gin.Context) {
	var language models.Language
	language_id, _ := strconv.Atoi(c.Param("id"))

	if err := models.DB.Where("id = ?", language_id).First(&language).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Language not found!"})
		return
	}

	models.DB.Delete(&language)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
