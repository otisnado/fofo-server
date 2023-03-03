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

	return Routes
}
