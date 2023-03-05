package main

import (
	"github.com/otisnado/fofo-server/models"
	"github.com/otisnado/fofo-server/routes"
)

func main() {
	models.ConnectDatabase()
	router := routes.InitRouter()
	router.Run()

}
