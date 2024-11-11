package server

import (
	"github.com/ebilsanta/social-network/backend/complex_services/follower_complex_service/controllers"
	follower "github.com/ebilsanta/social-network/backend/complex_services/follower_complex_service/controllers/followers"
	"github.com/ebilsanta/social-network/backend/complex_services/follower_complex_service/services"
	pb "github.com/ebilsanta/social-network/backend/complex_services/follower_complex_service/services/proto/generated"
	"github.com/gin-gonic/gin"
)

func NewRouter(followerClient pb.FollowerServiceClient, userClient pb.UserServiceClient, producer *services.KafkaProducer) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	followerController := follower.NewFollowerController(followerClient, userClient, producer)

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("api/v1")
	{
		followerGroup := v1.Group("followers")
		{
			followerGroup.GET("/:id", followerController.GetFollowers)
			followerGroup.POST("/", followerController.AddFollower)
			followerGroup.DELETE("/", followerController.DeleteFollower)
		}
		followingGroup := v1.Group("followings")
		{
			followingGroup.GET("/:id", followerController.GetFollowing)
		}
	}
	return router

}
