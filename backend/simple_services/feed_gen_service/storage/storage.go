package storage

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type Storage interface {
	AddToFeeds([]string, string) error
}

type RedisStore struct {
	Client *redis.Client
}

func NewRedisStore() (*RedisStore, error) {
	client, err := connectToDB()
	if err != nil {
		return nil, err
	}

	return &RedisStore{Client: client}, nil
}

func connectToDB() (*redis.Client, error) {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		return nil, err
	}
	return redis.NewClient(opt), nil
}

func (s *RedisStore) AddToFeeds(users []string, postID string) error {
	pipe := s.Client.Pipeline()
	ctx := context.Background()

	for _, user := range users {
		key := fmt.Sprintf("feed:%s", user)
		score := float64(time.Now().Unix())
		pipe.ZAdd(ctx, key, redis.Z{
			Score:  score,
			Member: postID,
		})
	}
	_, err := pipe.Exec(ctx)
	return err
}
