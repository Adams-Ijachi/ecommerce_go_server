package main

import (
	"ecom/config"
	"ecom/core"

	pg "ecom/datastore/postgres"

	"log"
)

func main() {

	configuration := config.Init()
	err := configuration.LoadEnvVariables()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connecting to a database")

	client := pg.NewClient(*configuration)

	if err != nil {
		log.Fatal(err)
	}

	client.DB.Migrator().DropTable(&core.User{}, &core.UserProfile{})

	log.Println("Dropping completed")

}
