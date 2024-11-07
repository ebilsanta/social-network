package user

import (
	"fmt"
	"net/http"

	"github.com/ebilsanta/social-network/backend/complex_services/user_service/models"
	pb "github.com/ebilsanta/social-network/backend/complex_services/user_service/services/proto/generated"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type UserController struct {
	client pb.UserServiceClient
}

func NewUserController(client pb.UserServiceClient) *UserController {
	return &UserController{
		client: client,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"details": err.Error()})
		return
	}

	createdUser, err := uc.client.CreateUser(ctx, &pb.CreateUserRequest{
		Username: user.Username,
		Email:    user.Email,
		Image:    user.Image,
		Id:       user.Id,
		Name:     user.Name,
	})

	if err != nil {
		grpcStatus := status.Code(err)
		if grpcStatus == codes.AlreadyExists {
			ctx.JSON(http.StatusConflict, gin.H{
				"error":   "failed to create user",
				"details": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user", "details": err.Error()})
		return
	}

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

	users, err := uc.client.GetUsers(ctx, &pb.GetUsersRequest{
		Query: ctx.Query("query"),
		Page:  params.Page,
		Limit: params.Limit,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get users", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, mapGetUsersResponse(users))
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := uc.client.GetUser(ctx, &pb.GetUserRequest{Id: id})
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
