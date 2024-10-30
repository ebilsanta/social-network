package storage

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/ebilsanta/social-network/backend/user-service/proto/generated"
	"github.com/ebilsanta/social-network/backend/user-service/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage interface {
	CreateUser(*types.User) (*pb.User, error)
	GetUsers() ([]*pb.User, error)
	GetUser(string) (*pb.User, error)
	DeleteUser(string) error
	UpdatePostCount(string, int32) error
	UpdateFollowerFollowingCount(string, string, int32) error
}

type MongoStore struct {
	Client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoStore() (*MongoStore, error) {
	client, err := connectToDB()
	if err != nil {
		return nil, err
	}
	collection := client.Database("user").Collection("user")
	return &MongoStore{Client: client, collection: collection}, nil
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

func (s *MongoStore) CreateUser(user *types.User) (*pb.User, error) {
	result, err := s.collection.InsertOne(context.TODO(), user)

	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:             result.InsertedID.(primitive.ObjectID).Hex(),
		Email:          user.Email,
		Username:       user.Username,
		ImageURL:       user.ImageURL,
		PostCount:      0,
		FollowerCount:  0,
		FollowingCount: 0,
		CreatedAt:      timestamppb.New(user.CreatedAt),
		DeletedAt:      nil,
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
		var user types.User
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

	var user types.User
	if err := s.collection.FindOne(context.TODO(), primitive.M{"_id": objID}).Decode(&user); err != nil {
		return nil, err
	}

	return decodeUser(user), nil
}

func (s *MongoStore) DeleteUser(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = s.collection.DeleteOne(context.TODO(), primitive.M{"_id": objID})
	return err
}

func decodeUser(user types.User) *pb.User {
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

func (s *MongoStore) UpdatePostCount(id string, change int32) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$inc": bson.M{"postCount": change}}

	_, err = s.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (s *MongoStore) UpdateFollowerFollowingCount(followerId, followingId string, change int32) error {
	followerObjID, err := primitive.ObjectIDFromHex(followerId)
	if err != nil {
		return err
	}

	followingObjID, err := primitive.ObjectIDFromHex(followingId)
	if err != nil {
		return err
	}

	followerFilter := bson.M{"_id": followerObjID}
	followingFilter := bson.M{"_id": followingObjID}

	followerUpdate := bson.M{"$inc": bson.M{"followingCount": change}}
	followingUpdate := bson.M{"$inc": bson.M{"followerCount": change}}

	session, err := s.Client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.Background())

	err = mongo.WithSession(context.Background(), session, func(sc mongo.SessionContext) error {
		if _, err := s.collection.UpdateOne(sc, followerFilter, followerUpdate); err != nil {
			return err
		}
		if _, err := s.collection.UpdateOne(sc, followingFilter, followingUpdate); err != nil {
			return err
		}
		return nil
	})

	return err
}
