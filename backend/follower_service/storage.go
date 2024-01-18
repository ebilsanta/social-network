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
	driver *neo4j.DriverWithContext
}

func NewGraphStore() (*GraphStore, error) {
	driver, err := connectToDB()
	if err != nil {
		return nil, err
	}

	return &GraphStore{driver: driver}, nil
}

func connectToDB() (*neo4j.DriverWithContext, error) {
	ctx := context.Background()
	dbUri := os.Getenv("DB_URI")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	fmt.Println("credentials:")
	fmt.Println(dbUri)
	fmt.Println(dbUser)
	fmt.Println(dbPassword)
	driver, _ := neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""))
	defer driver.Close(ctx)

	err := driver.VerifyConnectivity(ctx)
	if err != nil {
		return nil, err
	}
	return &driver, nil
}

func (s *GraphStore) GetFollowers(userID int) ([]*User, error) {
	return nil, nil
}

func (s *GraphStore) GetFollowing(userID int) ([]*User, error) {
	return nil, nil
}

func (s *GraphStore) AddFollower(followerID, followedUserID int) error {
	return nil
}

func (s *GraphStore) DeleteFollower(followerID, followedUserID int) error {
	return nil
}
