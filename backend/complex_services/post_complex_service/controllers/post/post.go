package user

import (
	"net/http"

	"github.com/ebilsanta/social-network/backend/complex_services/post_service/models"
	pb "github.com/ebilsanta/social-network/backend/complex_services/post_service/services/proto/generated"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type PostController struct {
	client pb.PostServiceClient
}

func NewPostController(client pb.PostServiceClient) *PostController {
	return &PostController{
		client: client,
	}
}

func (pc *PostController) CreatePost(ctx *gin.Context) {
	var post models.CreatePostRequest
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not parse body", "details": err.Error()})
		return
	}

	createdPost, err := pc.client.CreatePost(ctx, &pb.CreatePostRequest{
		Image:   post.Image,
		Caption: post.Caption,
		UserId:  post.UserId,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create post", "details": err.Error()})
		return
	}

	resp := &models.CreatePostResponse{
		Data: createdPost,
	}

	ctx.JSON(http.StatusCreated, resp)
}

func (pc *PostController) GetPostsByUserId(ctx *gin.Context) {
	id := ctx.Param("id")
	var query models.GetPostsByUserRequest
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse query", "details": err.Error()})
		return
	}

	posts, err := pc.client.GetPostsByUserId(ctx, &pb.GetPostsByUserRequest{
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
