package server

import (
	"github.com/ebilsanta/social-network/backend/complex_services/user_service/controllers"
	"github.com/ebilsanta/social-network/backend/complex_services/user_service/controllers/user"
	pb "github.com/ebilsanta/social-network/backend/complex_services/user_service/services/proto/generated"
	"github.com/gin-gonic/gin"
)

func NewRouter(userClient pb.UserServiceClient) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	userController := user.NewUserController(userClient)

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("v1")
	{
		usersGroup := v1.Group("users")
		{
			usersGroup.GET("/", userController.GetUsers)
			usersGroup.POST("/", userController.CreateUser)
			usersGroup.GET("/:id", userController.GetUser)
		}
	}
	return router

}
