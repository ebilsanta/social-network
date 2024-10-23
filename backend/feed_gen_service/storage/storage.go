package storage

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "github.com/ebilsanta/social-network/backend/feed-gen-service/api/proto/generated"
	"github.com/redis/go-redis/v9"
)

type Storage interface {
	AddToFeeds([]*pb.User, string) error
}

type RedisStore struct {
	Client *redis.Client
}

func NewRedisStore() (*RedisStore, error) {
	client, err := connectToDB()
	if err != nil {
		return nil, err
	}
	addMultiplePostsToFeed(client, "1")
	addMultiplePostsToFeed(client, "2")

	return &RedisStore{Client: client}, nil
}

func addMultiplePostsToFeed(rdb *redis.Client, userID string) error {
	posts := []string{"1", "2", "3", "4", "5"}
	ctx := context.Background()
	key := fmt.Sprintf("feed:%s", userID)
	pipe := rdb.Pipeline()

	for _, post := range posts {
		score := float64(time.Now().Unix())
		pipe.ZAdd(ctx, key, redis.Z{
			Score:  score,
			Member: post,
		})
	}
	_, err := pipe.Exec(ctx)
	return err
}

func connectToDB() (*redis.Client, error) {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		return nil, err
	}
	return redis.NewClient(opt), nil
}

func (s *RedisStore) AddToFeeds(users []*pb.User, postID string) error {
	pipe := s.Client.Pipeline()
	ctx := context.Background()

	for _, user := range users {
		key := fmt.Sprintf("feed:%s", user.Id)
		score := float64(time.Now().Unix())
		pipe.ZAdd(ctx, key, redis.Z{
			Score:  score,
			Member: postID,
		})
	}
	_, err := pipe.Exec(ctx)
	return err
}
