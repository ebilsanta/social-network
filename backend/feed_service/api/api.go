package api

import (
	"context"
	"fmt"
	"strconv"

	pb "github.com/ebilsanta/social-network/backend/feed-service/api/proto/generated"
	"github.com/ebilsanta/social-network/backend/feed-service/storage"
)

type FeedServiceServer struct {
	pb.UnimplementedFeedServiceServer
	store          storage.Storage
	followerClient pb.FollowerServiceClient
	postClient     pb.PostServiceClient
}

func NewServer(store storage.Storage, followerClient pb.FollowerServiceClient, postClient pb.PostServiceClient) *FeedServiceServer {
	return &FeedServiceServer{
		store:          store,
		followerClient: followerClient,
		postClient:     postClient,
	}
}

func (s *FeedServiceServer) GetFeed(ctx context.Context, req *pb.GetFeedRequest) (*pb.GetFeedResponse, error) {
	feed, err := s.store.GetFeed(req.UserId, req.Offset, req.Limit)

	if err != nil {
		return nil, err
	}

	if len(feed) != 0 && req.Offset == 0 {
		postIds := make([]int64, 0, len(feed))
		for _, post := range feed {
			postId, err := strconv.ParseInt(post, 10, 64)
			if err != nil {
				continue
			}
			postIds = append(postIds, postId)
		}
		posts, err := s.postClient.GetPostsByPostIds(ctx, &pb.GetPostsByIdsRequest{PostIds: postIds})
		if err != nil {
			return nil, err
		}
		return &pb.GetFeedResponse{Posts: posts.Posts}, nil
	}

	return s.createFeed(ctx, req)
}

func (s *FeedServiceServer) createFeed(ctx context.Context, req *pb.GetFeedRequest) (*pb.GetFeedResponse, error) {
	following, err := s.followerClient.GetFollowing(ctx, &pb.GetFollowingRequest{Id: req.UserId})
	if err != nil {
		return nil, err
	}
	userIds := make([]string, len(following.Following))

	for i, user := range following.Following {
		userIds[i] = user.Id
	}
	fmt.Printf("get feed from userIds: %v\n", userIds)
	posts, err := s.postClient.GetPostsByUserIds(ctx, &pb.GetPostsByUsersRequest{UserIds: userIds})

	if err != nil {
		return nil, err
	}
	return &pb.GetFeedResponse{Posts: posts.Posts}, nil
}
