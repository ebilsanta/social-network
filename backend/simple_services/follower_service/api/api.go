package api

import (
	"context"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	pb "github.com/ebilsanta/social-network/backend/follower-service/proto/generated"
	"github.com/ebilsanta/social-network/backend/follower-service/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FollowerServiceServer struct {
	pb.UnimplementedFollowerServiceServer
	store storage.Storage
	kafka *KafkaClient
}

func NewServer(store storage.Storage, kafka *KafkaClient) *FollowerServiceServer {
	return &FollowerServiceServer{
		store: store,
		kafka: kafka,
	}
}

func (s *FollowerServiceServer) AddUser(ctx context.Context, req *pb.AddUserRequest) (*emptypb.Empty, error) {
	err := s.store.AddUser(req.Id)
	return nil, err
}

func (s *FollowerServiceServer) GetFollowers(ctx context.Context, req *pb.GetFollowersRequest) (*pb.GetFollowersResponse, error) {
	if req.Limit <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "limit must be a positive integer")
	}
	return s.store.GetFollowers(req.Id, req.Page, req.Limit)
}

func (s *FollowerServiceServer) GetFollowing(ctx context.Context, req *pb.GetFollowingRequest) (*pb.GetFollowingResponse, error) {
	if req.Limit <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "limit must be a positive integer")
	}
	return s.store.GetFollowing(req.Id, req.Page, req.Limit)
}

func (s *FollowerServiceServer) AddFollower(ctx context.Context, req *pb.AddFollowerRequest) (*emptypb.Empty, error) {
	err := s.store.AddFollower(req.FollowerID, req.FollowedID)
	return nil, err
}

func (s *FollowerServiceServer) DeleteFollower(ctx context.Context, req *pb.DeleteFollowerRequest) (*emptypb.Empty, error) {
	err := s.store.DeleteFollower(req.FollowerID, req.FollowedID)
	return nil, err
}

func (s *FollowerServiceServer) ListenKafkaEvents(quit chan struct{}) {
	for {
		select {
		case <-quit:
			s.kafka.producer.Close()
			s.kafka.consumer.Close()
			log.Default().Println("Kafka producer and consumer closed.")
			return
		case e := <-s.kafka.producer.Events():
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					log.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		default:
			ev, err := s.kafka.consumer.ReadMessage(100 * time.Millisecond)
			if err == nil {
				userId := string(ev.Key)
				switch *ev.TopicPartition.Topic {
				case "new-user.add-graph-user":
					log.Default().Printf("Add user to graph: userId: %s\n", userId)
					s.store.AddUser(userId)
				}
			}
		}
	}
}
