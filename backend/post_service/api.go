package main

import (
	"context"
	"errors"

	pb "github.com/ebilsanta/social-network/backend/post-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PostServiceServer struct {
	pb.UnimplementedPostServiceServer
	store Storage
}

func newServer(store Storage) *PostServiceServer {
	return &PostServiceServer{
		store: store,
	}
}

func (s *PostServiceServer) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.Post, error) {
	post := NewPost(req.Caption, req.ImageURL, req.UserId)
	dbPost, err := s.store.CreatePost(post)

	if err != nil {
		return nil, HandleError(err)
	}

	return dbPost, nil
}

func (s *PostServiceServer) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.Post, error) {
	post, err := s.store.GetPostByID(req.Id)
	if err != nil {
		return nil, HandleError(err)
	}
	return post, nil
}

func (s *PostServiceServer) GetPosts(ctx context.Context, req *pb.GetPostsRequest) (*pb.GetPostsResponse, error) {
	posts, err := s.store.GetPosts()
	if err != nil {
		return nil, HandleError(err)
	}
	return &pb.GetPostsResponse{Posts: posts}, nil
}

func (s *PostServiceServer) GetPostByUserID(ctx context.Context, req *pb.GetPostByUserRequest) (*pb.GetPostsResponse, error) {
	posts, err := s.store.GetPostsByUserID(req.Id)
	if err != nil {
		return nil, HandleError(err)
	}
	return &pb.GetPostsResponse{Posts: posts}, nil
}

func HandleError(err error) error {
	var notFoundErr *PostNotFoundError
	if errors.As(err, &notFoundErr) {
		return status.Errorf(codes.NotFound, "post with id %d not found", notFoundErr.PostID)
	}

	var postgresErr *PostgresError
	if errors.As(err, &postgresErr) {
		return status.Errorf(codes.Internal, postgresErr.Error())
	}

	return status.Errorf(codes.Unknown, "an unexpected error occurred: %v", err)
}
