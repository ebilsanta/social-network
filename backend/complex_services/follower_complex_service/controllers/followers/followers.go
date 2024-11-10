package followers

import (
	"net/http"

	"github.com/ebilsanta/social-network/backend/complex_services/follower_complex_service/models"
	"github.com/ebilsanta/social-network/backend/complex_services/follower_complex_service/services"
	pb "github.com/ebilsanta/social-network/backend/complex_services/follower_complex_service/services/proto/generated"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type FollowerController struct {
	followerClient pb.FollowerServiceClient
	userClient     pb.UserServiceClient
	producer       *services.KafkaProducer
}

func NewFollowerController(followerClient pb.FollowerServiceClient, userClient pb.UserServiceClient, producer *services.KafkaProducer) *FollowerController {
	return &FollowerController{
		followerClient: followerClient,
		userClient:     userClient,
		producer:       producer,
	}
}

func (c *FollowerController) GetFollowers(ctx *gin.Context) {
	var params models.GetRequest
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not parse query params", "details": err.Error()})
		return
	}

	followerIds, err := c.followerClient.GetFollowers(ctx, &pb.GetFollowersRequest{
		Id:    ctx.Param("id"),
		Page:  int32(params.Page),
		Limit: int32(params.Limit),
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get followers", "details": err.Error()})
		return
	}

	if followerIds.Data == nil || len(followerIds.Data) == 0 {
		ctx.JSON(http.StatusOK, &models.GetResponse{
			Data:       []*pb.User{},
			Pagination: mapPaginationMetadata(followerIds.Pagination),
		})
		return
	}

	followers, err := c.userClient.GetUsersByIds(ctx, &pb.GetUsersByIdsRequest{
		Ids: followerIds.Data,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get followers", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &models.GetResponse{
		Data:       followers.Data,
		Pagination: mapPaginationMetadata(followerIds.Pagination),
	})
}

func (c *FollowerController) GetFollowing(ctx *gin.Context) {
	var params models.GetRequest
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not parse query params", "details": err.Error()})
		return
	}

	followingIds, err := c.followerClient.GetFollowing(ctx, &pb.GetFollowingRequest{
		Id:    ctx.Param("id"),
		Page:  int32(params.Page),
		Limit: int32(params.Limit),
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get following list", "details": err.Error()})
		return
	}

	if followingIds.Data == nil || len(followingIds.Data) == 0 {
		ctx.JSON(http.StatusOK, &models.GetResponse{
			Data:       []*pb.User{},
			Pagination: mapPaginationMetadata(followingIds.Pagination),
		})
		return
	}

	followings, err := c.userClient.GetUsersByIds(ctx, &pb.GetUsersByIdsRequest{
		Ids: followingIds.Data,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get following list", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &models.GetResponse{
		Data:       followings.Data,
		Pagination: mapPaginationMetadata(followingIds.Pagination),
	})
}

func (c *FollowerController) AddFollower(ctx *gin.Context) {
	var body models.ModifyRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not parse body", "details": err.Error()})
		return
	}

	_, err := c.followerClient.AddFollower(ctx, &pb.AddFollowerRequest{
		FollowerID: body.FollowerId,
		FollowedID: body.FollowingId,
	})

	if err != nil {
		if grpcStatus, ok := status.FromError(err); ok {
			if grpcStatus.Code() == codes.AlreadyExists {
				ctx.JSON(http.StatusConflict, gin.H{
					"error":   "failed to add follower",
					"details": grpcStatus.Message(),
				})
				return
			} else if grpcStatus.Code() == codes.InvalidArgument {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error":   "failed to add follower",
					"details": grpcStatus.Message(),
				})
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add follower", "details": err.Error()})
		return
	}

	followerId := []byte(body.FollowerId)
	followedId := []byte(body.FollowingId)
	c.producer.Produce("new-follower.update-profile", followerId, followedId)
	c.producer.Produce("new-follower.notification", followerId, followedId)

	ctx.JSON(http.StatusCreated, gin.H{"message": "follower added"})
}

func (c *FollowerController) DeleteFollower(ctx *gin.Context) {
	var body models.ModifyRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not parse body", "details": err.Error()})
		return
	}

	_, err := c.followerClient.DeleteFollower(ctx, &pb.DeleteFollowerRequest{
		FollowerID: body.FollowerId,
		FollowedID: body.FollowingId,
	})

	if err != nil {
		if grpcStatus, ok := status.FromError(err); ok {
			if grpcStatus.Code() == codes.NotFound {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error":   "failed to delete follower",
					"details": grpcStatus.Message(),
				})
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete follower", "details": err.Error()})
		return
	}

	followerId := []byte(body.FollowerId)
	followedId := []byte(body.FollowingId)
	c.producer.Produce("delete-follower.update-profile", followerId, followedId)

	ctx.JSON(http.StatusOK, gin.H{"message": "follower deleted"})
}

func mapPaginationMetadata(grpcPagination *pb.FollowerPaginationMetadata) *models.PaginationMetadata {
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
