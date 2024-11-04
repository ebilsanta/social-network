package server

import (
	"github.com/ebilsanta/social-network/backend/complex_services/user_service/controllers"
	"github.com/ebilsanta/social-network/backend/complex_services/user_service/controllers/user"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("v1")
	{
		usersGroup := v1.Group("user")
		{
			usersGroup.GET("/:id", user.CreateUser)
		}
	}
	return router

}
