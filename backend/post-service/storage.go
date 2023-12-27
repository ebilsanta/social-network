package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreatePost(*Post) (*Post, error)
	DeletePost(int) error
	UpdatePost(*Post) error
	GetPosts() ([]*Post, error)
	GetPostByID(int) (*Post, error)
	GetPostsByUserID(int) ([]*Post, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := os.Getenv("POSTGRES_CONN_STR")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) Init() error {
	return s.createPostTable()
}

func (s *PostgresStore) createPostTable() error {
	query := `CREATE TABLE if not exists post (
		id serial primary key,
		caption varchar(2000),
		image_url varchar(2000),
		user_id serial,
		created_at timestamp,
		deleted_at timestamp
	)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreatePost(post *Post) (*Post, error) {
	statement := `
		INSERT INTO post (caption, image_url, user_id, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := s.db.QueryRow(
		statement,
		post.Caption,
		post.ImageURL,
		post.UserID,
		post.CreatedAt,
	).Scan(&post.ID)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *PostgresStore) UpdatePost(post *Post) error {
	return nil
}

func (s *PostgresStore) DeletePost(id int) error {
	statement := "UPDATE post SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1"
	_, err := s.db.Query(statement, id)
	return err
}

func (s *PostgresStore) GetPosts() ([]*Post, error) {
	statement := "SELECT * FROM post WHERE deleted_at IS NULL"
	rows, err := s.db.Query(statement)

	if err != nil {
		return nil, err
	}
	posts := []*Post{}
	for rows.Next() {
		post, err := scanIntoPost(rows)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (s *PostgresStore) GetPostByID(id int) (*Post, error) {
	statement := "SELECT * FROM post WHERE id = $1 AND deleted_at IS NULL"
	rows, err := s.db.Query(statement, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoPost(rows)
	}
	return nil, fmt.Errorf("post %d not found", id)
}

func (s *PostgresStore) GetPostsByUserID(id int) ([]*Post, error) {
	statement := "SELECT * FROM post WHERE user_id = $1 AND deleted_at IS NULL"
	rows, err := s.db.Query(statement, id)
	if err != nil {
		return nil, err
	}
	posts := []*Post{}
	for rows.Next() {
		post, err := scanIntoPost(rows)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func scanIntoPost(rows *sql.Rows) (*Post, error) {
	post := Post{}
	err := rows.Scan(
		&post.ID,
		&post.Caption,
		&post.ImageURL,
		&post.UserID,
		&post.CreatedAt,
		&post.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &post, nil
}
