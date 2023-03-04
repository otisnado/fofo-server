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

	/* Users routes */

	// Find all users registered
	Routes.GET("/users", controllers.FindUsers)

	// Find a specific user by its id
	Routes.GET("/users/:id", controllers.FindUser)

	// Create a user
	Routes.POST("/users", controllers.CreateUser)

	// Update data for a user --> id is required
	Routes.PATCH("/users/:id", controllers.UpdateUser)

	// Delete a user --> id is required
	Routes.DELETE("/users/:id", controllers.DeleteUser)

	return Routes
}
