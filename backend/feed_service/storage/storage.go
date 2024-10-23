package storage

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type Storage interface {
	GetFeed(string, int32, int32) ([]string, error)
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

func (s *RedisStore) GetFeed(id string, offset, limit int32) ([]string, error) {
	key := fmt.Sprintf("feed:%s", id)

	now := time.Now().Unix()
	twoDaysAgo := now - 2*24*60*60

	posts, err := s.Client.ZRevRangeByScore(
		context.Background(),
		key,
		&redis.ZRangeBy{
			Min:    fmt.Sprintf("%d", twoDaysAgo),
			Max:    fmt.Sprintf("%d", now),
			Offset: int64(offset),
			Count:  int64(limit),
		},
	).Result()

	return posts, err
}
