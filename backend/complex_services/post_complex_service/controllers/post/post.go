package post

import (
	"net/http"
	"strconv"

	"github.com/ebilsanta/social-network/backend/complex_services/post_service/models"
	"github.com/ebilsanta/social-network/backend/complex_services/post_service/services"
	pb "github.com/ebilsanta/social-network/backend/complex_services/post_service/services/proto/generated"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type PostController struct {
	postClient pb.PostServiceClient
	userClient pb.UserServiceClient
	producer   *services.KafkaProducer
}

func NewPostController(postClient pb.PostServiceClient, userClient pb.UserServiceClient, producer *services.KafkaProducer) *PostController {
	return &PostController{
		postClient: postClient,
		userClient: userClient,
		producer:   producer,
	}
}

func (pc *PostController) CreatePost(ctx *gin.Context) {
	var post models.CreatePostRequest
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not parse body", "details": err.Error()})
		return
	}

	createdPost, err := pc.postClient.CreatePost(ctx, &pb.CreatePostRequest{
		Image:   post.Image,
		Caption: post.Caption,
		UserId:  post.UserId,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create post", "details": err.Error()})
		return
	}

	key := []byte(post.UserId)
	value := []byte(strconv.FormatInt(createdPost.Id, 10))

	pc.producer.Produce("new-post.notification", key, value)
	pc.producer.Produce("new-post.update-feed", key, value)
	pc.producer.Produce("new-post.update-profile", key, value)

	resp := &models.CreatePostResponse{
		Data: createdPost,
	}

	ctx.JSON(http.StatusCreated, resp)
}

func (pc *PostController) GetPostById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id", "details": err.Error()})
		return
	}
	post, err := pc.postClient.GetPost(ctx, &pb.GetPostRequest{
		Id: int64(id),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get post", "details": err.Error()})
		return
	}
	user, err := pc.userClient.GetUsersByIds(ctx, &pb.GetUsersByIdsRequest{
		Ids: []string{post.UserId},
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user", "details": err.Error()})
		return
	}
	if len(user.Data) == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}
	user.Data[0].Email = ""

	resp := &models.GetPostResponse{
		Data: &models.Post{
			Id:        post.Id,
			Caption:   post.Caption,
			Image:     post.Image,
			User:      user.Data[0],
			CreatedAt: post.CreatedAt,
			DeletedAt: post.DeletedAt,
		},
	}

	ctx.JSON(http.StatusOK, resp)
}

func (pc *PostController) GetPostsByUserId(ctx *gin.Context) {
	id := ctx.Param("id")
	var query models.GetPostsByUserRequest
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse query", "details": err.Error()})
		return
	}

	posts, err := pc.postClient.GetPostsByUserId(ctx, &pb.GetPostsByUserRequest{
		Id:    id,
		Page:  int32(query.Page),
		Limit: int32(query.Limit),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get posts", "details": err.Error()})
		return
	}

	if posts.Data == nil {
		posts.Data = []*pb.Post{}
	}

	resp := &models.GetPostsResponse{
		Data:       posts.Data,
		Pagination: mapPaginationMetadata(posts.Pagination),
	}

	ctx.JSON(http.StatusOK, resp)
}

func mapPaginationMetadata(grpcPagination *pb.PostPaginationMetadata) *models.PaginationMetadata {
	return &models.PaginationMetadata{
		TotalRecords: int64(grpcPagination.TotalRecords),
		CurrentPage:  int64(grpcPagination.CurrentPage),
		TotalPages:   int64(grpcPagination.TotalPages),
		NextPage:     intValueToInt(grpcPagination.NextPage),
		PrevPage:     intValueToInt(grpcPagination.PrevPage),
	}
}

func intValueToInt(wrapper *wrapperspb.Int32Value) *int64 {
	if wrapper != nil {
		val := int64(wrapper.Value)
		return &val
	}
	return nil
}
