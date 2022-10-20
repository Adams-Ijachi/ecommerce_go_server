package main

import (
	"ecom/config"
	"ecom/controllers"

	pg "ecom/datastore/postgres"
	"ecom/routes"
	"log"

	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
)

func main() {

	configuration := config.Init()
	err := configuration.LoadEnvVariables()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connecting to a database")

	database := pg.NewClient(*configuration)

	authController := controllers.NewAuthController(database)

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

	routes.InitializeAllRoutes(router, authController)

	router.Run()

}
