package api

import (
	"context"
	"errors"
	"log"

	"github.com/ebilsanta/social-network/backend/follower-service/errtypes"
	pb "github.com/ebilsanta/social-network/backend/follower-service/proto/generated"
	"github.com/ebilsanta/social-network/backend/follower-service/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FollowerServiceServer struct {
	pb.UnimplementedFollowerServiceServer
	store    storage.Storage
	producer *KafkaProducer
}

func NewServer(store storage.Storage, producer *KafkaProducer) *FollowerServiceServer {
	return &FollowerServiceServer{
		store:    store,
		producer: producer,
	}
}

func (s *FollowerServiceServer) AddUser(ctx context.Context, req *pb.AddUserRequest) (*emptypb.Empty, error) {
	log.Default().Printf("follower_service AddUser request: %v", req)
	err := s.store.AddUser(req.Id)
	if err != nil {
		return nil, HandleError(err)
	}
	return nil, nil
}

func (s *FollowerServiceServer) GetFollowers(ctx context.Context, req *pb.GetFollowersRequest) (*pb.GetFollowersResponse, error) {
	log.Default().Printf("follower_service GetFollowers request: %v", req)
	followers, err := s.store.GetFollowers(req.Id, req.Page, req.Limit)
	if err != nil {
		return nil, HandleError(err)
	}
	return followers, nil
}

func (s *FollowerServiceServer) GetFollowing(ctx context.Context, req *pb.GetFollowingRequest) (*pb.GetFollowingResponse, error) {
	log.Default().Printf("follower_service GetFollowing request: %v", req)
	following, err := s.store.GetFollowing(req.Id, req.Page, req.Limit)
	if err != nil {
		return nil, HandleError(err)
	}
	return following, nil
}

func (s *FollowerServiceServer) AddFollower(ctx context.Context, req *pb.AddFollowerRequest) (*emptypb.Empty, error) {
	log.Default().Printf("follower_service AddFollower request: %v", req)
	err := s.store.AddFollower(req.FollowerID, req.FollowedID)
	if err != nil {
		return nil, HandleError(err)
	}
	followerId := []byte(req.FollowerID)
	followedId := []byte(req.FollowedID)
	s.producer.Produce("new-follower.update-profile", followerId, followedId)
	s.producer.Produce("new-follower.notification", followerId, followedId)
	return nil, nil
}

func (s *FollowerServiceServer) DeleteFollower(ctx context.Context, req *pb.DeleteFollowerRequest) (*emptypb.Empty, error) {
	log.Default().Printf("follower_service DeleteFollower request: %v", req)
	err := s.store.DeleteFollower(req.FollowerID, req.FollowedID)
	if err != nil {
		return nil, HandleError(err)
	}
	followerId := []byte(req.FollowerID)
	followedId := []byte(req.FollowedID)
	s.producer.Produce("delete-follower.update-profile", followerId, followedId)
	return nil, nil
}

func HandleError(err error) error {
	var notFoundErr *errtypes.UserNotFoundError
	if errors.As(err, &notFoundErr) {
		return status.Errorf(codes.NotFound, "user with id %s not found", notFoundErr.UserID)
	}

	var neo4jErr *errtypes.Neo4jError
	if errors.As(err, &neo4jErr) {
		return status.Errorf(codes.Internal, "neo4j database error: %v", neo4jErr)
	}

	return status.Errorf(codes.Unknown, "an unexpected error occurred: %v", err)
}
