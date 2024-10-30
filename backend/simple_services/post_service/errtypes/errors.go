package errtypes

import "fmt"

type PostNotFoundError struct {
	PostID int64
}

type PostgresError struct {
	Query string
	Err   error
}

func (e *PostNotFoundError) Error() string {
	return fmt.Sprintf("post with id %d not found", e.PostID)
}

func NewPostNotFoundError(postID int64) error {
	return &PostNotFoundError{PostID: postID}
}

func (e *PostgresError) Error() string {
	return fmt.Sprintf("postgres query failed: %s, error: %v", e.Query, e.Err)
}

func NewPostgresError(query string, err error) error {
	return &PostgresError{
		Query: query,
		Err:   err,
	}
}
