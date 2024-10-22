package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/ebilsanta/social-network/backend/user-service/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage interface {
	CreateUser(*User) (*pb.User, error)
	GetUsers() ([]*pb.User, error)
	GetUser(string) (*pb.User, error)
}

type MongoStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoStore() (*MongoStore, error) {
	client, err := connectToDB()
	if err != nil {
		return nil, err
	}
	collection := client.Database("user").Collection("user")
	return &MongoStore{client: client, collection: collection}, nil
}

func connectToDB() (*mongo.Client, error) {
	uri := os.Getenv("MONGODB_URI")
	log.Default().Printf("USER_SERVICE Connecting to MongoDB at %s", uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return client, nil
}

func (s *MongoStore) CreateUser(user *User) (*pb.User, error) {
	result, err := s.collection.InsertOne(context.TODO(), user)

	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:        result.InsertedID.(primitive.ObjectID).Hex(),
		Email:     user.Email,
		Username:  user.Username,
		ImageURL:  user.ImageURL,
		CreatedAt: timestamppb.New(user.CreatedAt),
		DeletedAt: nil,
	}, nil
}

func (s *MongoStore) GetUsers() ([]*pb.User, error) {
	cursor, err := s.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var users []*pb.User
	for cursor.Next(context.Background()) {
		var user User
		if err = cursor.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, decodeUser(user))
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *MongoStore) GetUser(id string) (*pb.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user User
	if err := s.collection.FindOne(context.TODO(), primitive.M{"_id": objID}).Decode(&user); err != nil {
		return nil, err
	}

	return decodeUser(user), nil
}

func decodeUser(user User) *pb.User {
	var deletedAt *timestamppb.Timestamp
	if user.DeletedAt != nil {
		deletedAt = timestamppb.New(*user.DeletedAt)
	}

	return &pb.User{
		Id:        user.Id.Hex(),
		Email:     user.Email,
		Username:  user.Username,
		ImageURL:  user.ImageURL,
		CreatedAt: timestamppb.New(user.CreatedAt),
		DeletedAt: deletedAt,
	}
}
