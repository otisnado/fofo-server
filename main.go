package main

import (
	"github.com/otisnado/fofo-server/models"
	"github.com/otisnado/fofo-server/routes"
)

func main() {
	models.ConnectDatabase()
	routes.Routes().Run()

}
