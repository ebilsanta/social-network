package feed

import (
	"net/http"

	"github.com/ebilsanta/social-network/backend/complex_services/feed_complex_service/models"
	pb "github.com/ebilsanta/social-network/backend/complex_services/feed_complex_service/services/proto/generated"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type FeedController struct {
	feedClient pb.FeedServiceClient
	postClient pb.PostServiceClient
	userClient pb.UserServiceClient
}

func NewFeedController(feedClient pb.FeedServiceClient, postClient pb.PostServiceClient, userClient pb.UserServiceClient) *FeedController {
	return &FeedController{
		feedClient: feedClient,
		postClient: postClient,
		userClient: userClient,
	}
}

func (c *FeedController) GetFeed(ctx *gin.Context) {
	var params models.GetFeedRequest
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not parse query params", "details": err.Error()})
		return
	}

	postIds, err := c.feedClient.GetFeed(ctx, &pb.GetFeedRequest{
		UserId: ctx.Param("id"),
		Page:   int32(params.Page),
		Limit:  int32(params.Limit),
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get post ids for feed", "details": err.Error()})
		return
	}

	if len(postIds.Data) == 0 {
		ctx.JSON(http.StatusOK, &models.GetFeedResponse{
			Data:       []*models.Post{},
			Pagination: mapPaginationMetadata(postIds.Pagination),
		})
		return
	}

	posts, err := c.postClient.GetPostsByPostIds(ctx, &pb.GetPostsByIdsRequest{
		PostIds: postIds.Data,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get feed", "details": err.Error()})
		return
	}

	userIds := make([]string, 0, len(posts.Posts))
	for _, post := range posts.Posts {
		userIds = append(userIds, post.UserId)
	}

	users, err := c.userClient.GetUsersByIds(ctx, &pb.GetUsersByIdsRequest{Ids: userIds})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get users", "details": err.Error()})
		return
	}

	userMap := make(map[string]*pb.User)
	for _, user := range users.Data {
		userMap[user.Id] = user
		userMap[user.Email] = nil
	}

	postsWithUser := make([]*models.Post, 0, len(posts.Posts))
	for _, post := range posts.Posts {
		postsWithUser = append(postsWithUser, &models.Post{
			Id:        post.Id,
			Caption:   post.Caption,
			Image:     post.Image,
			User:      userMap[post.UserId],
			CreatedAt: post.CreatedAt,
			DeletedAt: post.DeletedAt,
		})
	}

	ctx.JSON(http.StatusOK, &models.GetFeedResponse{
		Data:       postsWithUser,
		Pagination: mapPaginationMetadata(postIds.Pagination),
	})
}

func mapPaginationMetadata(grpcPagination *pb.FeedPaginationMetadata) *models.PaginationMetadata {
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
