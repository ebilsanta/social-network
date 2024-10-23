package storage

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ebilsanta/social-network/backend/post-service/errtypes"
	pb "github.com/ebilsanta/social-network/backend/post-service/proto"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreatePost(*pb.Post) (*pb.Post, error)
	GetPosts() ([]*pb.Post, error)
	GetPostById(int64) (*pb.Post, error)
	GetPostsByPostIds([]int64) ([]*pb.Post, error)
	GetPostsByUserId(string) ([]*pb.Post, error)
	GetPostsByUserIds([]string) ([]*pb.Post, error)
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
		post.UserId,
		post.CreatedAt.AsTime(),
	).Scan(&post.Id)

	if err != nil {
		return nil, errtypes.NewPostgresError(statement, err)
	}

	return post, nil
}

func (s *PostgresStore) GetPosts() ([]*pb.Post, error) {
	statement := "SELECT * FROM post WHERE deleted_at IS NULL"
	rows, err := s.db.Query(statement)

	if err != nil {
		return nil, errtypes.NewPostgresError(statement, err)
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

func (s *PostgresStore) GetPostById(id int64) (*pb.Post, error) {
	statement := "SELECT * FROM post WHERE id = $1 AND deleted_at IS NULL"
	rows, err := s.db.Query(statement, id)
	if err != nil {
		return nil, errtypes.NewPostgresError(statement, err)
	}
	for rows.Next() {
		return scanIntoPost(rows)
	}
	return nil, errtypes.NewPostNotFoundError(id)
}

func (s *PostgresStore) GetPostsByPostIds(postIds []int64) ([]*pb.Post, error) {
	if len(postIds) == 0 {
		return []*pb.Post{}, nil
	}

	placeholders := make([]string, len(postIds))
	for i := range postIds {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	placeholdersList := strings.Join(placeholders, ",")

	statement := fmt.Sprintf(`
		SELECT * FROM post
		WHERE id IN (%s) 
		AND deleted_at IS NULL
		ORDER BY created_at DESC`, placeholdersList)

	params := make([]interface{}, len(postIds))
	for i, id := range postIds {
		params[i] = id
	}

	rows, err := s.db.Query(statement, params...)
	if err != nil {
		return nil, errtypes.NewPostgresError(statement, err)
	}
	defer rows.Close()

	posts := []*pb.Post{}
	for rows.Next() {
		post, err := scanIntoPost(rows)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *PostgresStore) GetPostsByUserId(id string) ([]*pb.Post, error) {
	statement := "SELECT * FROM post WHERE user_id = $1 AND deleted_at IS NULL"
	rows, err := s.db.Query(statement, id)
	if err != nil {
		return nil, errtypes.NewPostgresError(statement, err)
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

func (s *PostgresStore) GetPostsByUserIds(userIds []string) ([]*pb.Post, error) {
	if len(userIds) == 0 {
		return []*pb.Post{}, nil
	}

	placeholders := make([]string, len(userIds))
	for i := range userIds {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	placeholdersList := strings.Join(placeholders, ",")

	statement := fmt.Sprintf(`
		SELECT * FROM post 
		WHERE user_id IN (%s) AND deleted_at IS NULL 
		ORDER BY created_at DESC 
		LIMIT 10`, placeholdersList)
	params := make([]interface{}, len(userIds))
	for i, id := range userIds {
		params[i] = id
	}

	rows, err := s.db.Query(statement, params...)
	if err != nil {
		return nil, errtypes.NewPostgresError(statement, err)
	}
	defer rows.Close()

	posts := []*pb.Post{}
	for rows.Next() {
		post, err := scanIntoPost(rows)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
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
		&post.UserId,
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
