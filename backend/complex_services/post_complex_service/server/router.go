package server

import (
	"github.com/ebilsanta/social-network/backend/complex_services/post_service/controllers"
	post "github.com/ebilsanta/social-network/backend/complex_services/post_service/controllers/post"
	"github.com/ebilsanta/social-network/backend/complex_services/post_service/services"
	pb "github.com/ebilsanta/social-network/backend/complex_services/post_service/services/proto/generated"
	"github.com/gin-gonic/gin"
)

func NewRouter(postClient pb.PostServiceClient, producer *services.KafkaProducer) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	postController := post.NewPostController(postClient, producer)

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("api/v1")
	{
		postsGroup := v1.Group("posts")
		{
			postsGroup.POST("/", postController.CreatePost)
			postsGroup.GET("/user/:id", postController.GetPostsByUserId)
		}
	}
	return router

}
