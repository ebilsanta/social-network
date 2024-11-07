package storage

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "github.com/ebilsanta/social-network/backend/user-service/proto/generated"
	"github.com/ebilsanta/social-network/backend/user-service/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage interface {
	CreateUser(*types.User) (*pb.User, error)
	GetUsers(string, int64, int64) (*pb.GetUsersResponse, error)
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

	// Enable text searches on username field
	usernameIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: "text"}},
		Options: options.Index().SetUnique(true),
	}

	idIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "id", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err = collection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{usernameIndex, idIndex})
	if err != nil {
		return nil, err
	}

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
	_, err := s.collection.InsertOne(context.TODO(), user)

	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:             user.Id,
		Email:          user.Email,
		Name:           user.Name,
		Username:       user.Username,
		Image:          user.Image,
		PostCount:      0,
		FollowerCount:  0,
		FollowingCount: 0,
		CreatedAt:      timestamppb.New(user.CreatedAt),
		DeletedAt:      nil,
	}, nil
}

func (s *MongoStore) GetUsers(query string, page, limit int64) (*pb.GetUsersResponse, error) {
	filter := bson.M{
		"username": bson.M{
			"$regex":   query,
			"$options": "i",
		},
	}
	totalRecords, err := s.collection.CountDocuments(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	totalPages := (totalRecords + limit - 1) / limit
	skip := (page - 1) * limit

	findOpt := options.Find().SetSkip(skip).SetLimit(limit)
	cursor, err := s.collection.Find(context.TODO(), filter, findOpt)
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

	var nextPage, prevPage *wrapperspb.Int32Value
	if page < totalPages {
		nextPage = &wrapperspb.Int32Value{Value: int32(page + 1)}
	}
	if page > 1 {
		prevPage = &wrapperspb.Int32Value{Value: int32(page - 1)}
	}

	return &pb.GetUsersResponse{
		Data: users,
		Pagination: &pb.UserPaginationMetadata{
			TotalRecords: totalRecords,
			CurrentPage:  page,
			TotalPages:   totalPages,
			NextPage:     nextPage,
			PrevPage:     prevPage,
		},
	}, nil
}

func (s *MongoStore) GetUser(id string) (*pb.User, error) {
	var user types.User
	if err := s.collection.FindOne(context.TODO(), primitive.M{"id": id}).Decode(&user); err != nil {
		return nil, err
	}

	return decodeUser(user), nil
}

func (s *MongoStore) DeleteUser(id string) error {
	_, err := s.collection.DeleteOne(context.TODO(), primitive.M{"id": id})
	return err
}

func decodeUser(user types.User) *pb.User {
	var deletedAt *timestamppb.Timestamp
	if user.DeletedAt != nil {
		deletedAt = timestamppb.New(*user.DeletedAt)
	}

	return &pb.User{
		Id:             user.Id,
		Email:          user.Email,
		Name:           user.Name,
		Username:       user.Username,
		Image:          user.Image,
		PostCount:      user.PostCount,
		FollowerCount:  user.FollowerCount,
		FollowingCount: user.FollowingCount,
		CreatedAt:      timestamppb.New(user.CreatedAt),
		DeletedAt:      deletedAt,
	}
}

func (s *MongoStore) UpdatePostCount(id string, change int32) error {
	filter := bson.M{"id": id}
	update := bson.M{"$inc": bson.M{"postCount": change}}

	_, err := s.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (s *MongoStore) UpdateFollowerFollowingCount(followerId, followingId string, change int32) error {
	followerFilter := bson.M{"id": followerId}
	followingFilter := bson.M{"id": followerId}

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
