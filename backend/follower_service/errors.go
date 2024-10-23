package main

import "fmt"

type UserNotFoundError struct {
	UserID string
}

type Neo4jError struct {
	Query string
	Err   error
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("user with id %s not found", e.UserID)
}

func NewUserNotFoundError(userID string) error {
	return &UserNotFoundError{UserID: userID}
}

func (e *Neo4jError) Error() string {
	return fmt.Sprintf("neo4j query failed: %s, error: %v", e.Query, e.Err)
}

func NewNeo4jError(query string, err error) error {
	return &Neo4jError{
		Query: query,
		Err:   err,
	}
}
