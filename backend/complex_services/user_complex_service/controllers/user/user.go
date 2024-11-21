package user

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ebilsanta/social-network/backend/complex_services/user_service/models"
	"github.com/ebilsanta/social-network/backend/complex_services/user_service/services"
	pb "github.com/ebilsanta/social-network/backend/complex_services/user_service/services/proto/generated"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type UserController struct {
	userClient     pb.UserServiceClient
	followerClient pb.FollowerServiceClient
	producer       *services.KafkaProducer
}

func NewUserController(userClient pb.UserServiceClient, followerClient pb.FollowerServiceClient, producer *services.KafkaProducer) *UserController {
	return &UserController{
		userClient:     userClient,
		followerClient: followerClient,
		producer:       producer,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not parse body", "details": err.Error()})
		return
	}

	createdUser, err := uc.userClient.CreateUser(ctx, &pb.CreateUserRequest{
		Username: user.Username,
		Email:    user.Email,
		Image:    user.Image,
		Id:       user.Id,
		Name:     user.Name,
	})

	if err != nil {
		grpcStatus := status.Code(err)
		if grpcStatus == codes.AlreadyExists {
			details := "user with this id, email, or username already exists"
			duplicateFields := map[string]string{
				"email":    "user with this email already exists",
				"username": "user with this username already exists",
				"id":       "user with this id already exists",
			}
			for field, msg := range duplicateFields {
				if strings.Contains(err.Error(), field) {
					details = msg
					break
				}
			}
			ctx.JSON(http.StatusConflict, gin.H{"error": "failed to create user", "details": details})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user", "details": err.Error()})
		return
	}

	key := []byte(createdUser.Id)
	uc.producer.Produce("new-user.add-graph-user", key, []byte(""))

	resp := &models.CreateUserResponse{
		Data: &models.User{
			Id:             createdUser.Id,
			Email:          createdUser.Email,
			Name:           createdUser.Name,
			Username:       createdUser.Username,
			Image:          createdUser.Image,
			PostCount:      0,
			FollowerCount:  0,
			FollowingCount: 0,
			CreatedAt:      createdUser.CreatedAt,
		},
	}

	ctx.JSON(http.StatusCreated, resp)
}

func (uc *UserController) GetUsers(ctx *gin.Context) {
	var params models.GetUsersRequest
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse params", "details": err.Error()})
		return
	}

	users, err := uc.userClient.GetUsers(ctx, &pb.GetUsersRequest{
		Query: params.Query,
		Page:  params.Page,
		Limit: params.Limit,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get users", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, mapGetUsersResponse(users))
}

func (uc *UserController) GetIdHandler(c *gin.Context) {
	followedId := c.Param("followedId")
	if followedId == "" {
		uc.GetUser(c)
	} else {
		uc.CheckFollowing(c)
	}
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := uc.userClient.GetUser(ctx, &pb.GetUserRequest{Id: id})
	if err != nil {
		grpcStatus := status.Code(err)
		if grpcStatus == codes.NotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   "user not found",
				"details": fmt.Sprintf("user with id %s not found", id),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &models.GetUserResponse{
		Data: mapUser(user, true),
	})
}

func (uc *UserController) GetUserByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := uc.userClient.GetUserByUsername(ctx, &pb.GetUserByUsernameRequest{Username: username})
	if err != nil {
		grpcStatus := status.Code(err)
		if grpcStatus == codes.NotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   "user not found",
				"details": fmt.Sprintf("user with username %s not found", username),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &models.GetUserResponse{
		Data: mapUser(user, false),
	})
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not parse body", "details": err.Error()})
		return
	}

	updatedUser, err := uc.userClient.UpdateUser(ctx, &pb.UpdateUserRequest{
		Id:       id,
		Username: user.Username,
		Email:    user.Email,
		Image:    user.Image,
		Name:     user.Name,
	})

	if err != nil {
		if grpcStatus, ok := status.FromError(err); ok {
			switch grpcStatus.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, gin.H{
					"error":   "user not found",
					"details": grpcStatus.Message(),
				})
				return
			case codes.AlreadyExists:
				details := "user with this id, email, or username already exists"
				duplicateFields := map[string]string{
					"email":    "user with this email already exists",
					"username": "user with this username already exists",
					"id":       "user with this id already exists",
				}
				for field, msg := range duplicateFields {
					if strings.Contains(grpcStatus.Message(), field) {
						details = msg
						break
					}
				}
				ctx.JSON(http.StatusConflict, gin.H{
					"error":   "failed to update user",
					"details": details,
				})
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &models.GetUserResponse{
		Data: mapUser(updatedUser, true),
	})
}

func (uc *UserController) CheckFollowing(ctx *gin.Context) {
	followerID := ctx.Param("id")
	followedID := ctx.Param("followedId")
	resp, err := uc.followerClient.CheckFollowing(ctx, &pb.CheckFollowingRequest{
		FollowerID: followerID,
		FollowedID: followedID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check following", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &models.CheckFollowingResponse{
		Following: resp.Following,
	})
}

func mapUser(pbUser *pb.User, showEmail bool) *models.User {
	user := &models.User{
		Id:             pbUser.Id,
		Name:           pbUser.Name,
		Username:       pbUser.Username,
		PostCount:      pbUser.PostCount,
		FollowerCount:  pbUser.FollowerCount,
		FollowingCount: pbUser.FollowingCount,
		Image:          pbUser.Image,
		CreatedAt:      pbUser.CreatedAt,
		DeletedAt:      pbUser.DeletedAt,
	}
	if showEmail {
		user.Email = pbUser.Email
	}
	return user
}

func mapGetUsersResponse(grpcResponse *pb.GetUsersResponse) *models.GetUsersResponse {
	users := []*models.User{}
	for _, grpcUser := range grpcResponse.Data {
		users = append(users, mapUser(grpcUser, false))
	}

	pagination := &models.PaginationMetadata{}
	if grpcResponse.Pagination != nil {
		pagination = &models.PaginationMetadata{
			TotalRecords: grpcResponse.Pagination.TotalRecords,
			CurrentPage:  grpcResponse.Pagination.CurrentPage,
			TotalPages:   grpcResponse.Pagination.TotalPages,
			NextPage:     intValueToInt(grpcResponse.Pagination.NextPage),
			PrevPage:     intValueToInt(grpcResponse.Pagination.PrevPage),
		}
	}

	return &models.GetUsersResponse{
		Data:       users,
		Pagination: pagination,
	}
}

func intValueToInt(wrapper *wrapperspb.Int32Value) *int64 {
	if wrapper != nil {
		val := int64(wrapper.Value)
		return &val
	}
	return nil
}
