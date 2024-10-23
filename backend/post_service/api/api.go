package api

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/ebilsanta/social-network/backend/post-service/errtypes"
	pb "github.com/ebilsanta/social-network/backend/post-service/proto"
	"github.com/ebilsanta/social-network/backend/post-service/storage"
	"github.com/ebilsanta/social-network/backend/post-service/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PostServiceServer struct {
	pb.UnimplementedPostServiceServer
	store    storage.Storage
	producer *KafkaProducer
}

func NewServer(store storage.Storage, producer *KafkaProducer) *PostServiceServer {
	return &PostServiceServer{
		store:    store,
		producer: producer,
	}
}

func (s *PostServiceServer) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.Post, error) {
	log.Default().Printf("post_service CreatePost request: %v", req)
	post := types.NewPost(req.Caption, req.ImageURL, req.UserId)
	dbPost, err := s.store.CreatePost(post)

	if err != nil {
		return nil, HandleError(err)
	}

	key := []byte(req.UserId)
	value := []byte(strconv.FormatInt(dbPost.Id, 10))

	s.producer.Produce(key, value)

	return dbPost, nil
}

func (s *PostServiceServer) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.Post, error) {
	log.Default().Printf("post_service GetPost request: %v", req)
	post, err := s.store.GetPostById(req.Id)
	if err != nil {
		return nil, HandleError(err)
	}
	return post, nil
}

func (s *PostServiceServer) GetPosts(ctx context.Context, req *pb.GetPostsRequest) (*pb.GetPostsResponse, error) {
	log.Default().Printf("post_service GetPosts request: %v", req)
	posts, err := s.store.GetPosts()
	if err != nil {
		return nil, HandleError(err)
	}
	return &pb.GetPostsResponse{Posts: posts}, nil
}

func (s *PostServiceServer) GetPostsByPostIds(ctx context.Context, req *pb.GetPostsByIdsRequest) (*pb.GetPostsResponse, error) {
	log.Default().Printf("post_service GetPostsByPostIds request: %v", req)
	posts, err := s.store.GetPostsByPostIds(req.PostIds)
	if err != nil {
		return nil, HandleError(err)
	}
	return &pb.GetPostsResponse{Posts: posts}, nil
}

func (s *PostServiceServer) GetPostsByUserId(ctx context.Context, req *pb.GetPostsByUserRequest) (*pb.GetPostsResponse, error) {
	log.Default().Printf("post_service GetPostsByUserId request: %v", req)
	posts, err := s.store.GetPostsByUserId(req.Id)
	if err != nil {
		return nil, HandleError(err)
	}
	return &pb.GetPostsResponse{Posts: posts}, nil
}

func (s *PostServiceServer) GetPostsByUserIds(ctx context.Context, req *pb.GetPostsByUsersRequest) (*pb.GetPostsResponse, error) {
	log.Default().Printf("post_service GetPostsByUserIds request: %v", req)
	posts, err := s.store.GetPostsByUserIds(req.UserIds, req.Offset, req.Limit)
	if err != nil {
		return nil, HandleError(err)
	}
	return &pb.GetPostsResponse{Posts: posts}, nil
}

func HandleError(err error) error {
	var notFoundErr *errtypes.PostNotFoundError
	if errors.As(err, &notFoundErr) {
		return status.Errorf(codes.NotFound, "post with id %d not found", notFoundErr.PostID)
	}

	var postgresErr *errtypes.PostgresError
	if errors.As(err, &postgresErr) {
		return status.Errorf(codes.Internal, postgresErr.Error())
	}

	return status.Errorf(codes.Unknown, "an unexpected error occurred: %v", err)
}
