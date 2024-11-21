package storage

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	pb "github.com/ebilsanta/social-network/backend/feed-service/proto/generated"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Storage interface {
	GetFeed(string, int32, int32) (*pb.GetFeedResponse, error)
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

func (s *RedisStore) GetFeed(id string, page, limit int32) (*pb.GetFeedResponse, error) {
	key := fmt.Sprintf("feed:%s", id)

	now := time.Now().Unix()
	sevenDaysAgo := now - 7*24*60*60
	offset := (page - 1) * limit

	posts, err := s.Client.ZRevRangeByScore(
		context.Background(),
		key,
		&redis.ZRangeBy{
			// Min:    fmt.Sprintf("%d", sevenDaysAgo),
			Min:    "0",
			Max:    fmt.Sprintf("%d", now),
			Offset: int64(offset),
			Count:  int64(limit),
		},
	).Result()

	if err != nil {
		return nil, err
	}

	postIds := make([]int64, 0, len(posts))
	for _, post := range posts {
		postId, err := strconv.ParseInt(post, 10, 64)
		if err != nil {
			continue
		}
		postIds = append(postIds, postId)
	}

	totalRecords, err := s.Client.ZCount(
		context.Background(),
		key,
		fmt.Sprintf("%d", sevenDaysAgo),
		fmt.Sprintf("%d", now),
	).Result()
	if err != nil {
		return nil, err
	}

	totalPages := int32(totalRecords) / limit
	if totalRecords%int64(limit) > 0 {
		totalPages++
	}

	var nextPage, prevPage *wrapperspb.Int32Value
	if page < totalPages {
		nextPage = &wrapperspb.Int32Value{Value: page + 1}
	}
	if page > 1 {
		prevPage = &wrapperspb.Int32Value{Value: page - 1}
	}

	return &pb.GetFeedResponse{
		Data: postIds,
		Pagination: &pb.FeedPaginationMetadata{
			TotalRecords: totalRecords,
			CurrentPage:  page,
			TotalPages:   totalPages,
			NextPage:     nextPage,
			PrevPage:     prevPage,
		},
	}, nil
}
