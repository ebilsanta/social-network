package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/ebilsanta/social-network/backend/follower-service/errtypes"
	pb "github.com/ebilsanta/social-network/backend/follower-service/proto/generated"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Storage interface {
	AddUser(string) error
	GetFollowers(string) ([]*pb.GraphUser, error)
	GetFollowing(string) ([]*pb.GraphUser, error)
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
		// "MERGE (p:User {id:'1'})",
		// "MERGE (p:User {id:'2'})",
		// "MERGE (p:User {id:'3'})",
		// "MATCH (user1:User {id: '1'}), (user2:User {id: '2'}) MERGE (user1)-[:FOLLOWS]->(user2)",
		// "MATCH (user1:User {id: '1'}), (user2:User {id: '3'}) MERGE (user1)-[:FOLLOWS]->(user2)",
		// "MATCH (user1:User {id: '2'}), (user2:User {id: '3'}) MERGE (user1)-[:FOLLOWS]->(user2)",
		// "MATCH (user1:User {id: '3'}), (user2:User {id: '2'}) MERGE (user1)-[:FOLLOWS]->(user2)",
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

func (s *GraphStore) GetFollowers(userID string) ([]*pb.GraphUser, error) {
	exists, err := s.UserExists(userID)
	if err != nil {
		return nil, errtypes.NewNeo4jError("checking if user exists", err)
	}
	if !exists {
		return nil, errtypes.NewUserNotFoundError(userID)
	}

	ctx := context.Background()
	query := `MATCH (:User {id: $id})<-[:FOLLOWS]-(follower)
			RETURN follower.id`
	result, err := neo4j.ExecuteQuery(ctx, s.Driver,
		query,
		map[string]any{
			"id": userID,
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"),
	)
	if err != nil {
		return nil, err
	}
	followers := []*pb.GraphUser{}
	for _, record := range result.Records {
		id, ok := record.Get("follower.id")
		if !ok {
			continue
		}
		followerID, ok := id.(string)
		if ok {
			follower := &pb.GraphUser{Id: followerID}
			followers = append(followers, follower)
		}
	}

	return followers, nil
}

func (s *GraphStore) GetFollowing(userID string) ([]*pb.GraphUser, error) {
	exists, err := s.UserExists(userID)
	if err != nil {
		return nil, errtypes.NewNeo4jError("checking if user exists", err)
	}
	if !exists {
		return nil, errtypes.NewUserNotFoundError(userID)
	}

	ctx := context.Background()
	query := `MATCH (:User {id: $id})-[:FOLLOWS]->(followed)
			RETURN followed.id`
	result, err := neo4j.ExecuteQuery(ctx, s.Driver,
		query,
		map[string]any{
			"id": userID,
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"),
	)
	if err != nil {
		return nil, err
	}
	followings := []*pb.GraphUser{}
	for _, record := range result.Records {
		id, ok := record.Get("followed.id")
		if !ok {
			continue
		}
		followerID, ok := id.(string)
		if ok {
			following := &pb.GraphUser{Id: followerID}
			followings = append(followings, following)
		}
	}

	return followings, nil
}

func (s *GraphStore) AddFollower(followerID, followedID string) error {
	if followerID == followedID {
		return fmt.Errorf("user %s cannot follow themselves", followerID)
	}

	exists, err := s.UserExists(followerID)
	if err != nil {
		return errtypes.NewNeo4jError("checking if follower exists", err)
	}
	if !exists {
		return errtypes.NewUserNotFoundError(followerID)
	}

	exists, err = s.UserExists(followedID)
	if err != nil {
		return err
	}
	if !exists {
		return errtypes.NewUserNotFoundError(followedID)
	}

	exists, err = s.UserFollows(followerID, followedID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("user %s already follows user %s", followerID, followedID)
	}

	ctx := context.Background()
	query := `MATCH (follower:User {id: $followerID}), (followed:User {id: $followedID}) 
	          RETURN follower, followed`
	params := map[string]any{
		"followerID": followerID,
		"followedID": followedID,
	}
	query = `MATCH (follower:User {id: $followerID}), (followed:User {id: $followedID}) 
			MERGE (follower)-[:FOLLOWS]->(followed)`
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

func (s *GraphStore) DeleteFollower(followerID, followedID string) error {
	exists, err := s.UserExists(followerID)
	if err != nil {
		return errtypes.NewNeo4jError("checking if follower exists", err)
	}
	if !exists {
		return errtypes.NewUserNotFoundError(followerID)
	}

	exists, err = s.UserExists(followedID)
	if err != nil {
		return errtypes.NewNeo4jError("checking if followed exists", err)
	}
	if !exists {
		return errtypes.NewUserNotFoundError(followerID)
	}

	exists, err = s.UserFollows(followerID, followedID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("user %s does not follow user %s", followerID, followedID)
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

func (s *GraphStore) UserExists(userID string) (bool, error) {
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

func (s *GraphStore) UserFollows(followerID, followedID string) (bool, error) {
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