package main

import (
	"github.com/otisnado/nepackage/models"
	"github.com/otisnado/nepackage/routes"
	"github.com/otisnado/nepackage/utils"
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
