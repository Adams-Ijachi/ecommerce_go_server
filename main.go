package main

import (
	"ecom/config"
	pg "ecom/datastore/postgres"
	"ecom/routes"

	"log"

	"github.com/gin-gonic/gin"
)

//TODO connect to db postgres
// TODO: create a .env file
// TODO: create a router
// TODO: create a request header check

func main() {

	configuration := config.Init()
	err := configuration.LoadEnvVariables()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connecting to a database")

	pg.NewClient(*configuration)

	router := gin.New()
	router.Use(config.Cors())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.InitializeAllRoutes(router)

	router.Run()

}
