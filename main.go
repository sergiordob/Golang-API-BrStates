package main

import (
	"github.com/sergiordob/Golang-API-REST-01/config"
	"github.com/sergiordob/Golang-API-REST-01/routes"
	"github.com/sergiordob/Golang-API-REST-01/database"
)

func main() {
	
	//init routes
	routes.LoadRoutes(config.Mux)

	database.OpenDB()

	//init server
	config.InitServer()
}
