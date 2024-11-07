package api

import (
	"context"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	pb "github.com/ebilsanta/social-network/backend/user-service/proto/generated"
	"github.com/ebilsanta/social-network/backend/user-service/storage"
	"github.com/ebilsanta/social-network/backend/user-service/types"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	store          storage.Storage
	followerClient pb.FollowerServiceClient
	consumer       *kafka.Consumer
}

func NewServer(store storage.Storage, followerClient pb.FollowerServiceClient, consumer *kafka.Consumer) *UserServiceServer {
	return &UserServiceServer{
		store:          store,
		followerClient: followerClient,
		consumer:       consumer,
	}
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	user := types.NewUser(req.Id, req.Email, req.Name, req.Username, req.Image)
	dbUser, err := s.store.CreateUser(user)

	if err != nil {
		return nil, err
	}

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

func (s *UserServiceServer) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	res, err := s.store.GetUsers(req.Query, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserServiceServer) StartUsersListener(quit chan struct{}) {
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
					key, val := string(ev.Key), string(ev.Value)
					switch *ev.TopicPartition.Topic {
					case "new-post.update-profile":
						userId, postId := key, val
						s.handleNewPosts(userId, postId)
					case "new-follower.update-profile":
						followerId, followingId := key, val
						s.handleNewFollower(followerId, followingId)
					case "delete-follower.update-profile":
						followerId, followingId := key, val
						s.handleDeleteFollower(followerId, followingId)
					}
				}
			}
		}
	}()
}

func (s *UserServiceServer) handleNewPosts(userId string, postId string) {
	log.Default().Printf("Update profile new post: userId: %s, postId: %s\n", userId, postId)
	err := s.store.UpdatePostCount(userId, 1)
	if err != nil {
		log.Default().Printf("Error updating post count: %v\n", err)
	}
}

func (s *UserServiceServer) handleNewFollower(followerId string, followingId string) {
	log.Default().Printf("Update profile new follower: followerId: %s, followingId: %s\n", followerId, followingId)
	err := s.store.UpdateFollowerFollowingCount(followerId, followingId, 1)
	if err != nil {
		log.Default().Printf("Error updating follower/following count: %v\n", err)
	}
}

func (s *UserServiceServer) handleDeleteFollower(followerId string, followingId string) {
	log.Default().Printf("Update profile delete follower: followerId: %s, followingId: %s\n", followerId, followingId)
	err := s.store.UpdateFollowerFollowingCount(followerId, followingId, -1)
	if err != nil {
		log.Default().Printf("Error updating follower/following count: %v\n", err)
	}
}
