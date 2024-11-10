package storage

import (
	"context"
	"fmt"
	"os"

	pb "github.com/ebilsanta/social-network/backend/follower-service/proto/generated"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Storage interface {
	AddUser(string) error
	GetFollowers(string, int32, int32) (*pb.GetFollowersResponse, error)
	GetFollowing(string, int32, int32) (*pb.GetFollowingResponse, error)
	AddFollower(string, string) error
	DeleteFollower(string, string) error
}

type GraphStore struct {
	Driver neo4j.DriverWithContext
	Ctx    context.Context
}

func NewGraphStore() (*GraphStore, error) {
	ctx := context.Background()
	driver, err := connectToDB(ctx)
	if err != nil {
		return nil, err
	}

	return &GraphStore{Driver: driver, Ctx: ctx}, nil
}

func connectToDB(ctx context.Context) (neo4j.DriverWithContext, error) {
	dbUri := os.Getenv("DB_URI")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	driver, _ := neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""))

	err := driver.VerifyConnectivity(ctx)
	if err != nil {
		return nil, err
	}
	return driver, nil
}

func (s *GraphStore) Init() error {
	ctx := context.Background()
	queries := []string{
		"CREATE CONSTRAINT IF NOT EXISTS FOR (u:User) REQUIRE (u.id) IS UNIQUE",
	}
	for _, query := range queries {
		_, err := neo4j.ExecuteQuery(ctx, s.Driver,
			query,
			nil,
			neo4j.EagerResultTransformer,
			neo4j.ExecuteQueryWithDatabase("neo4j"))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *GraphStore) AddUser(userID string) error {
	ctx := context.Background()
	query := `MERGE (u:User {id: $id}) RETURN u`
	_, err := neo4j.ExecuteQuery(ctx, s.Driver,
		query,
		map[string]any{
			"id": userID,
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"),
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *GraphStore) GetFollowers(userID string, page, limit int32) (*pb.GetFollowersResponse, error) {
	exists, err := s.userExists(userID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !exists {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("user with id %s not found", userID))
	}

	skip := (page - 1) * limit

	ctx := context.Background()
	query := `MATCH (:User {id: $id})<-[:FOLLOWS]-(follower)
			RETURN follower.id
			SKIP $skip LIMIT $limit`
	result, err := neo4j.ExecuteQuery(ctx, s.Driver,
		query,
		map[string]any{
			"id":    userID,
			"skip":  skip,
			"limit": limit,
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"),
	)
	if err != nil {
		return nil, err
	}

	followers := []string{}
	for _, record := range result.Records {
		id, ok := record.Get("follower.id")
		if !ok {
			continue
		}
		followerID, ok := id.(string)
		if ok {
			followers = append(followers, followerID)
		}
	}

	countQuery := `MATCH (:User {id: $id})<-[:FOLLOWS]-(follower)
			RETURN count(follower) as count`
	countResult, err := neo4j.ExecuteQuery(ctx, s.Driver,
		countQuery,
		map[string]any{
			"id": userID,
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"),
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	totalRecordsAny, _ := countResult.Records[0].Get("count")
	totalRecordsInt64 := totalRecordsAny.(int64)
	totalRecords := int32(totalRecordsInt64)
	totalPages := totalRecords / limit
	if (totalRecords % limit) > 0 {
		totalPages++
	}

	var nextPage, prevPage *wrapperspb.Int32Value
	if page < totalPages {
		nextPage = &wrapperspb.Int32Value{Value: page + 1}
	}
	if page > 1 {
		prevPage = &wrapperspb.Int32Value{Value: page - 1}
	}

	return &pb.GetFollowersResponse{
		Data: followers,
		Pagination: &pb.FollowerPaginationMetadata{
			TotalRecords: totalRecords,
			CurrentPage:  page,
			TotalPages:   totalPages,
			NextPage:     nextPage,
			PrevPage:     prevPage,
		},
	}, nil
}

func (s *GraphStore) GetFollowing(userID string, page, limit int32) (*pb.GetFollowingResponse, error) {
	exists, err := s.userExists(userID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !exists {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("user with id %s not found", userID))
	}

	skip := (page - 1) * limit

	ctx := context.Background()
	query := `MATCH (:User {id: $id})-[:FOLLOWS]->(followed)
			RETURN followed.id
			SKIP $skip LIMIT $limit`
	result, err := neo4j.ExecuteQuery(ctx, s.Driver,
		query,
		map[string]any{
			"id":    userID,
			"skip":  skip,
			"limit": limit,
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"),
	)
	if err != nil {
		return nil, err
	}
	followings := []string{}
	for _, record := range result.Records {
		id, ok := record.Get("followed.id")
		if !ok {
			continue
		}
		followerID, ok := id.(string)
		if ok {
			followings = append(followings, followerID)
		}
	}

	countQuery := `MATCH (:User {id: $id})-[:FOLLOWS]->(followed)
			RETURN count(followed) as count`
	countResult, err := neo4j.ExecuteQuery(ctx, s.Driver,
		countQuery,
		map[string]any{
			"id": userID,
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	totalRecordsAny, _ := countResult.Records[0].Get("count")
	totalRecordsInt64 := totalRecordsAny.(int64)
	totalRecords := int32(totalRecordsInt64)
	totalPages := totalRecords / limit
	if (totalRecords % limit) > 0 {
		totalPages++
	}

	var nextPage, prevPage *wrapperspb.Int32Value
	if page < totalPages {
		nextPage = &wrapperspb.Int32Value{Value: page + 1}
	}
	if page > 1 {
		prevPage = &wrapperspb.Int32Value{Value: page - 1}
	}

	return &pb.GetFollowingResponse{
		Data: followings,
		Pagination: &pb.FollowerPaginationMetadata{
			TotalRecords: totalRecords,
			CurrentPage:  page,
			TotalPages:   totalPages,
			NextPage:     nextPage,
			PrevPage:     prevPage,
		},
	}, nil
}

// Create users if they don't exist. If there is an existing relationship, return error.
func (s *GraphStore) AddFollower(followerID, followedID string) error {
	if followerID == followedID {
		return status.Errorf(codes.InvalidArgument, "user %s cannot follow themselves", followerID)
	}

	exists, err := s.userExists(followerID)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	if !exists {
		err = s.AddUser(followerID)
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}

	exists, err = s.userExists(followedID)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	if !exists {
		err = s.AddUser(followedID)
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}

	exists, err = s.userFollows(followerID, followedID)
	if err != nil {
		return err
	}
	if exists {
		return status.Error(codes.AlreadyExists, fmt.Sprintf("user %s already follows user %s", followerID, followedID))
	}

	ctx := context.Background()
	query := `MATCH (follower:User {id: $followerID}), (followed:User {id: $followedID})
              MERGE (follower)-[:FOLLOWS]->(followed)`
	params := map[string]any{
		"followerID": followerID,
		"followedID": followedID,
	}

	_, err = neo4j.ExecuteQuery(ctx, s.Driver,
		query,
		params,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"),
	)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func (s *GraphStore) DeleteFollower(followerID, followedID string) error {
	exists, err := s.userExists(followerID)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	if !exists {
		return status.Error(codes.NotFound, fmt.Sprintf("user %s not found", followerID))
	}

	exists, err = s.userExists(followedID)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	if !exists {
		return status.Error(codes.NotFound, fmt.Sprintf("user %s not found", followedID))
	}

	exists, err = s.userFollows(followerID, followedID)
	if err != nil {
		return err
	}
	if !exists {
		return status.Error(codes.NotFound, fmt.Sprintf("user %s does not follow user %s", followerID, followedID))
	}

	ctx := context.Background()
	query := `MATCH (follower:User {id: $followerID})-[r:FOLLOWS]->(followed:User {id: $followedID})
			DELETE r`
	params := map[string]any{
		"followerID": followerID,
		"followedID": followedID,
	}
	_, err = neo4j.ExecuteQuery(ctx, s.Driver,
		query,
		params,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"),
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *GraphStore) userExists(userID string) (bool, error) {
	ctx := context.Background()
	query := `MATCH (u:User {id: $id}) RETURN u`
	result, err := neo4j.ExecuteQuery(ctx, s.Driver,
		query,
		map[string]any{
			"id": userID,
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"),
	)
	if err != nil {
		return false, err
	}

	if len(result.Records) == 0 {
		return false, nil
	}

	return true, nil
}

func (s *GraphStore) userFollows(followerID, followedID string) (bool, error) {
	ctx := context.Background()
	checkQuery := `MATCH (follower:User {id: $followerID})-[r:FOLLOWS]->(followed:User {id: $followedID}) 
	               RETURN r`
	checkParams := map[string]any{
		"followerID": followerID,
		"followedID": followedID,
	}

	result, err := neo4j.ExecuteQuery(ctx, s.Driver,
		checkQuery,
		checkParams,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"),
	)
	if err != nil {
		return false, err
	}

	return len(result.Records) > 0, nil
}
