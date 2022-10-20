package routes

import (
	"ecom/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeAllRoutes(
	router *gin.Engine,
	authController *controllers.UserAuthController,

) {

	routeV1 := router.Group("/api/v1")

	userAuth := routeV1.Group("/auth")

	userAuth.Use()
	{
		userAuth.POST("/register", authController.RegisterUser())
	}

}
