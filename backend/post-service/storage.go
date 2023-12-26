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
		poster_id serial,
		created_at timestamp
	)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreatePost(post *Post) (*Post, error) {
	fmt.Printf("%+v\n", post)
	statement := `
		INSERT INTO post (caption, image_url, poster_id, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := s.db.QueryRow(
		statement,
		post.Caption,
		post.ImageURL,
		post.PosterID,
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
	return nil
}

func (s *PostgresStore) GetPosts() ([]*Post, error) {
	statement := "SELECT * FROM post"
	rows, err := s.db.Query(statement)

	if err != nil {
		return nil, err
	}
	posts := []*Post{}
	for rows.Next() {
		post := Post{}
		err := rows.Scan(
			&post.ID,
			&post.Caption,
			&post.ImageURL,
			&post.PosterID,
			&post.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (s *PostgresStore) GetPostByID(id int) (*Post, error) {
	return nil, nil
}
