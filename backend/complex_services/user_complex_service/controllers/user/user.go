package user

import (
	"net/http"

	"github.com/ebilsanta/social-network/backend/complex_services/user_service/models"
	pb "github.com/ebilsanta/social-network/backend/complex_services/user_service/services/proto/generated"
	"github.com/gin-gonic/gin"
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user", "details": err.Error()})
		return
	}

	// Protoc generates these fields as omitempty, so we need to set them manually
	createdUser.FollowerCount = 0
	createdUser.FollowingCount = 0
	createdUser.PostCount = 0

	ctx.JSON(http.StatusCreated, createdUser)
}
