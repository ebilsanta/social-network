package main

import (
	"context"

	pb "github.com/ebilsanta/social-network/backend/post-service/proto"
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
		return nil, err
	}

	return dbPost, nil
}

func (s *PostServiceServer) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.Post, error) {
	post, err := s.store.GetPostByID(req.Id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostServiceServer) GetPosts(ctx context.Context, req *pb.GetPostsRequest) (*pb.GetPostsResponse, error) {
	posts, err := s.store.GetPosts()
	if err != nil {
		return nil, err
	}
	return &pb.GetPostsResponse{Posts: posts}, nil
}

func (s *PostServiceServer) GetPostByUserID(ctx context.Context, req *pb.GetPostByUserRequest) (*pb.GetPostsResponse, error) {
	posts, err := s.store.GetPostsByUserID(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetPostsResponse{Posts: posts}, nil
}
