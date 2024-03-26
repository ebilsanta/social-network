package main

import (
	"context"
	"fmt"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Storage interface {
	GetFollowers(int) ([]*User, error)
	GetFollowing(int) ([]*User, error)
	AddFollower(int, int) error
	DeleteFollower(int, int) error
}

type GraphStore struct {
	driver neo4j.DriverWithContext
	ctx    context.Context
}

func NewGraphStore() (*GraphStore, error) {
	ctx := context.Background()
	driver, err := connectToDB(ctx)
	if err != nil {
		return nil, err
	}

	return &GraphStore{driver: driver, ctx: ctx}, nil
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
		"MERGE (p:User {id:1})",
		"MERGE (p:User {id:2})",
		"MERGE (p:User {id:3})",
		"MERGE (p:User {id:4})",
		"MATCH (user1:User {id: 1}), (user2:User {id: 2}) CREATE (user1)-[:FOLLOWS]->(user2)",
		"MATCH (user1:User {id: 1}), (user2:User {id: 3}) CREATE (user1)-[:FOLLOWS]->(user2)",
		"MATCH (user1:User {id: 2}), (user2:User {id: 1}) CREATE (user1)-[:FOLLOWS]->(user2)",
		"MATCH (user1:User {id: 2}), (user2:User {id: 3}) CREATE (user1)-[:FOLLOWS]->(user2)",
		"MATCH (user1:User {id: 3}), (user2:User {id: 2}) CREATE (user1)-[:FOLLOWS]->(user2)",
		"MATCH (user1:User {id: 3}), (user2:User {id: 1}) CREATE (user1)-[:FOLLOWS]->(user2)",
		"MATCH (user1:User {id: 3}), (user2:User {id: 4}) CREATE (user1)-[:FOLLOWS]->(user2)",
	}
	for _, query := range queries {
		_, err := neo4j.ExecuteQuery(ctx, s.driver,
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

func (s *GraphStore) GetFollowers(userID int) ([]*User, error) {
	ctx := context.Background()
	query := `MATCH (:User {id: $id})<--(follower)
			RETURN follower.id`
	result, err := neo4j.ExecuteQuery(ctx, s.driver,
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
	followers := []*User{}
	for _, record := range result.Records {
		fmt.Println(record.AsMap())
	}

	return followers, nil
}

func (s *GraphStore) GetFollowing(userID int) ([]*User, error) {
	return nil, nil
}

func (s *GraphStore) AddFollower(followerID, followedUserID int) error {
	// ctx := context.Background()
	// query := ""
	// result, _ := neo4j.ExecuteQuery(ctx, s.driver, )

	return nil
}

func (s *GraphStore) DeleteFollower(followerID, followedUserID int) error {
	return nil
}
