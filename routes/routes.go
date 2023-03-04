package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/otisnado/fofo-server/controllers"
)

func Routes() *gin.Engine {
	Routes := gin.Default()

	// Find all projects
	Routes.GET("/projects", controllers.FindProjects)

	// Find a specific project --> id is required
	Routes.GET("/projects/:id", controllers.FindProject)

	// Create a project
	Routes.POST("/projects", controllers.CreateProject)

	// Update data for a project --> id is required
	Routes.PATCH("/projects/:id", controllers.UpdateProject)

	// Delete a project --> id is required
	Routes.DELETE("/projects/:id", controllers.DeleteProject)

	/* Languages routes */

	// Find all languages supported
	Routes.GET("/languages", controllers.FindLanguages)

	// Find a specific language by its id
	Routes.GET("languages/:id", controllers.FindLanguage)

	// Create a language
	Routes.POST("/languages", controllers.CreateLanguage)

	// Update data for a language --> id is required
	Routes.PATCH("/languages/:id", controllers.UpdateLanguage)

	return Routes
}
