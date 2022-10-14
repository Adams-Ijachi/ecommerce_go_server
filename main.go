package main

import (
	"ecom/config"
	pg "ecom/datastore/postgres"
	"ecom/routes"

	"log"

	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
)

//TODO Create a models for db
// TODO: Create Services for db

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
	router.Use(secure.New(secure.Config{
		// AllowedHosts:          []string{"example.com", "ssl.example.com"},
		// SSLRedirect:           true,
		// SSLHost:               "ssl.example.com",
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
		IENoOpen:              true,
		ReferrerPolicy:        "strict-origin-when-cross-origin",
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
	}))

	routes.InitializeAllRoutes(router)

	router.Run()

}
