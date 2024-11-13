package api

import (
	"context"
	"log"

	pb "github.com/ebilsanta/social-network/backend/feed-service/proto/generated"
	"github.com/ebilsanta/social-network/backend/feed-service/storage"
)

type FeedServiceServer struct {
	pb.UnimplementedFeedServiceServer
	store storage.Storage
}

func NewServer(store storage.Storage) *FeedServiceServer {
	return &FeedServiceServer{
		store: store,
	}
}

func (s *FeedServiceServer) GetFeed(ctx context.Context, req *pb.GetFeedRequest) (*pb.GetFeedResponse, error) {
	log.Default().Printf("feed_service GetFeed request: %v", req)
	feed, err := s.store.GetFeed(req.UserId, req.Page, req.Limit)

	return feed, err
}
