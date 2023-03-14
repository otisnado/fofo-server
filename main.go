package main

import (
	"github.com/otisnado/fofo-server/models"
	"github.com/otisnado/fofo-server/routes"
	"github.com/otisnado/fofo-server/utils"
)

func main() {
	models.ConnectDatabase()
	utils.RoleAdminCreation()
	utils.PolicyAdminCreation()
	utils.GroupAdminCreation()
	utils.AdminCreation()
	router := routes.InitRouter()
	router.Run()

}
