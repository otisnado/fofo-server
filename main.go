package main

import (
	"github.com/otisnado/nepackage/models"
	"github.com/otisnado/nepackage/routes"
	"github.com/otisnado/nepackage/utils"
	log "github.com/sirupsen/logrus"
)

func main() {
	models.ConnectDatabase()
	utils.RoleAdminCreation()
	utils.PolicyAdminCreation()
	utils.GroupAdminCreation()
	utils.AdminCreation()
	router := routes.InitRouter()
	log.SetFormatter(&log.JSONFormatter{})
	router.Run()

}
