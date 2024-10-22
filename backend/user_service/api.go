package main

import (
	"context"

	pb "github.com/ebilsanta/social-network/backend/user-service/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	store Storage
}

func newServer(store Storage) *UserServiceServer {
	return &UserServiceServer{
		store: store,
	}
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	user := NewUser(req.Email, req.Username, req.ImageURL)
	dbUser, err := s.store.CreateUser(user)

	if err != nil {
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
