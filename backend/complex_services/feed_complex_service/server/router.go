package server

import (
	"github.com/ebilsanta/social-network/backend/complex_services/feed_complex_service/controllers"
	feed "github.com/ebilsanta/social-network/backend/complex_services/feed_complex_service/controllers/feed"
	pb "github.com/ebilsanta/social-network/backend/complex_services/feed_complex_service/services/proto/generated"
	"github.com/gin-gonic/gin"
)

func NewRouter(feedClient pb.FeedServiceClient, postClient pb.PostServiceClient, userClient pb.UserServiceClient) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	feedController := feed.NewFeedController(feedClient, postClient, userClient)

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("api/v1")
	{
		feedGroup := v1.Group("feeds")
		{
			feedGroup.GET("/:id", feedController.GetFeed)
		}
	}
	return router

}
