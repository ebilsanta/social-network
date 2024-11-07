package server

import (
	"github.com/ebilsanta/social-network/backend/complex_services/post_service/controllers"
	user "github.com/ebilsanta/social-network/backend/complex_services/post_service/controllers/post"
	pb "github.com/ebilsanta/social-network/backend/complex_services/post_service/services/proto/generated"
	"github.com/gin-gonic/gin"
)

func NewRouter(postClient pb.PostServiceClient) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	postController := user.NewPostController(postClient)

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("v1")
	{
		postsGroup := v1.Group("posts")
		{
			postsGroup.POST("/", postController.CreatePost)
			postsGroup.GET("/user/:id", postController.GetPostsByUserId)
		}
	}
	return router

}
