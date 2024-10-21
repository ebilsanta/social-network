package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/ebilsanta/social-network/backend/post-service/proto"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreatePost(*pb.Post) (*pb.Post, error)
	GetPosts() ([]*pb.Post, error)
	GetPostByID(int64) (*pb.Post, error)
	GetPostsByUserID(int64) ([]*pb.Post, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	db, err := connectToDB()
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil
}

func connectToDB() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", dbUser, dbName, dbPassword, dbHost, dbPort)
	return sql.Open("postgres", connStr)
}

func (s *PostgresStore) Init() error {
	return s.CreatePostTable()
}

func (s *PostgresStore) CreatePostTable() error {
	query := `CREATE TABLE if not exists post (
		id serial primary key,
		caption varchar(2000),
		image_url varchar(2000),
		user_id string,
		created_at timestamp,
		deleted_at timestamp
	)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreatePost(post *pb.Post) (*pb.Post, error) {
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
		post.CreatedAt.AsTime(),
	).Scan(&post.Id)

	if err != nil {
		return nil, NewPostgresError(statement, err)
	}

	return post, nil
}

func (s *PostgresStore) GetPosts() ([]*pb.Post, error) {
	statement := "SELECT * FROM post WHERE deleted_at IS NULL"
	rows, err := s.db.Query(statement)

	if err != nil {
		return nil, NewPostgresError(statement, err)
	}
	posts := []*pb.Post{}
	for rows.Next() {
		post, err := scanIntoPost(rows)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (s *PostgresStore) GetPostByID(id int64) (*pb.Post, error) {
	statement := "SELECT * FROM post WHERE id = $1 AND deleted_at IS NULL"
	rows, err := s.db.Query(statement, id)
	if err != nil {
		return nil, NewPostgresError(statement, err)
	}
	for rows.Next() {
		return scanIntoPost(rows)
	}
	return nil, NewPostNotFoundError(id)
}

func (s *PostgresStore) GetPostsByUserID(id int64) ([]*pb.Post, error) {
	statement := "SELECT * FROM post WHERE user_id = $1 AND deleted_at IS NULL"
	rows, err := s.db.Query(statement, id)
	if err != nil {
		return nil, NewPostgresError(statement, err)
	}
	posts := []*pb.Post{}
	for rows.Next() {
		post, err := scanIntoPost(rows)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func scanIntoPost(rows *sql.Rows) (*pb.Post, error) {
	var createdAt time.Time
	var deletedAt sql.NullTime

	post := pb.Post{}
	err := rows.Scan(
		&post.Id,
		&post.Caption,
		&post.ImageURL,
		&post.UserID,
		&createdAt,
		&deletedAt,
	)
	if err != nil {
		return nil, err
	}

	post.CreatedAt = timestamppb.New(createdAt)
	if deletedAt.Valid {
		post.DeletedAt = timestamppb.New(deletedAt.Time)
	} else {
		post.DeletedAt = nil
	}

	return &post, nil
}
