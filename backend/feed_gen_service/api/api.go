package api

import (
	"context"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	pb "github.com/ebilsanta/social-network/backend/feed-gen-service/api/proto/generated"
	"github.com/ebilsanta/social-network/backend/feed-gen-service/storage"
)

type FeedGenServiceServer struct {
	pb.UnimplementedFeedServiceServer
	store          storage.Storage
	followerClient pb.FollowerServiceClient
	consumer       *kafka.Consumer
}

func NewServer(store storage.Storage, followerClient pb.FollowerServiceClient, postClient pb.PostServiceClient, consumer *kafka.Consumer) *FeedGenServiceServer {
	return &FeedGenServiceServer{
		store:          store,
		followerClient: followerClient,
		consumer:       consumer,
	}
}

func (s *FeedGenServiceServer) StartPostsListener(quit chan struct{}) {
	go func() {
		for {
			select {
			case <-quit:
				s.consumer.Close()
				log.Println("Kafka consumer closed.")
				return
			default:
				ev, err := s.consumer.ReadMessage(100 * time.Millisecond)
				if err == nil {
					userId, postId := string(ev.Key), string(ev.Value)
					log.Default().Printf("Consumed message: userId: %s, postId: %s\n", userId, postId)
					s.updateFeeds(userId, postId)
				}
			}
		}
	}()
}

func (s *FeedGenServiceServer) updateFeeds(posterId, postId string) {
	followers, err := s.followerClient.GetFollowers(context.Background(), &pb.GetFollowersRequest{Id: posterId})
	if err != nil {
		log.Printf("Failed to get followers for user %s: %v\n", posterId, err)
		return
	}
	log.Printf("Adding post %s to feeds of %v followers\n", postId, followers.Followers)
	err = s.store.AddToFeeds(followers.Followers, postId)
	if err != nil {
		log.Printf("Failed to add post %s to feeds: %v\n", postId, err)
	}
}
