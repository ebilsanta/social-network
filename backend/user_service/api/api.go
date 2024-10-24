package api

import (
	"context"
	"log"

	pb "github.com/ebilsanta/social-network/backend/user-service/proto/generated"
	"github.com/ebilsanta/social-network/backend/user-service/storage"
	"github.com/ebilsanta/social-network/backend/user-service/types"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	store          storage.Storage
	followerClient pb.FollowerServiceClient
}

func NewServer(store storage.Storage, followerClient pb.FollowerServiceClient) *UserServiceServer {
	return &UserServiceServer{
		store:          store,
		followerClient: followerClient,
	}
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	user := types.NewUser(req.Email, req.Username, req.ImageURL)
	log.Default().Printf("user_service CreateUser request: %v", user)
	dbUser, err := s.store.CreateUser(user)

	if err != nil {
		return nil, err
	}
	log.Default().Printf("user_service CreateUser response: %v", dbUser)
	_, err = s.followerClient.AddUser(ctx, &pb.AddUserRequest{Id: dbUser.Id})
	if err != nil {
		s.store.DeleteUser(dbUser.Id)
		return nil, err
	}

	return dbUser, nil
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	user, err := s.store.GetUser(req.Id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceServer) GetUsers(ctx context.Context, req *emptypb.Empty) (*pb.GetUsersResponse, error) {
	users, err := s.store.GetUsers()
	if err != nil {
		return nil, err
	}
	return &pb.GetUsersResponse{Users: users}, nil
}
