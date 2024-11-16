package storage

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	GetUsersByIds([]string) (*pb.GetUsersByIdsResponse, error)
	GetUser(string) (*pb.User, error)
	GetUserByUsername(string) (*pb.User, error)
	UpdateUser(string, *string, *string, *string, *string) (*pb.User, error)
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

	emailIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err = collection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{usernameIndex, idIndex, emailIndex})
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
	var existingUser types.User
	err := s.collection.FindOne(context.TODO(), bson.M{"id": user.Id}).Decode(&existingUser)
	if err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "user with id %s already exists", user.Id)
	}

	// Check for unique email and username
	filter := bson.M{
		"$or": []bson.M{
			{"email": user.Email},
			{"username": user.Username},
		},
	}

	cursor, err := s.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, status.Error(codes.Internal, "error checking unique constraints")
	}
	defer cursor.Close(context.TODO())

	// Check if any other user has the same email or username
	for cursor.Next(context.TODO()) {
		var otherUser types.User
		if err := cursor.Decode(&otherUser); err == nil {
			if otherUser.Email == user.Email {
				return nil, status.Errorf(codes.AlreadyExists, "email %s already in use", user.Email)
			}
			if otherUser.Username == user.Username {
				return nil, status.Errorf(codes.AlreadyExists, "username %s already in use", user.Username)
			}
		}
	}

	_, err = s.collection.InsertOne(context.TODO(), user)
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
	err := s.collection.FindOne(context.TODO(), primitive.M{"id": id}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "user with id %s not found", id)
		}
		return nil, err
	}

	return decodeUser(user), nil
}

func (s *MongoStore) GetUserByUsername(username string) (*pb.User, error) {
	var user types.User
	err := s.collection.FindOne(context.TODO(), primitive.M{"username": username}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "user with username %s not found", username)
		}
		return nil, err
	}

	return decodeUser(user), nil
}

func (s *MongoStore) GetUsersByIds(ids []string) (*pb.GetUsersByIdsResponse, error) {
	filter := bson.M{"id": bson.M{"$in": ids}}
	log.Printf("ids: %v", ids)

	cursor, err := s.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var users []*pb.User
	for cursor.Next(context.TODO()) {
		var user types.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, decodeUser(user))
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	log.Printf("users: %v", users)

	return &pb.GetUsersByIdsResponse{
		Data: users,
	}, nil
}

func (s *MongoStore) UpdateUser(id string, email, name, username, image *string) (*pb.User, error) {
	// Check if the user with the given ID exists
	var existingUser types.User
	err := s.collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&existingUser)
	if err == mongo.ErrNoDocuments {
		return nil, status.Errorf(codes.NotFound, "user with id %s not found", id)
	} else if err != nil {
		return nil, status.Error(codes.Internal, "error retrieving user")
	}

	// Check for unique email and username if they are provided
	if email != nil || username != nil {
		filter := bson.M{
			"$or": []bson.M{},
			"id":  bson.M{"$ne": id},
		}
		if email != nil {
			filter["$or"] = append(filter["$or"].([]bson.M), bson.M{"email": *email})
		}
		if username != nil {
			filter["$or"] = append(filter["$or"].([]bson.M), bson.M{"username": *username})
		}

		cursor, err := s.collection.Find(context.TODO(), filter)
		if err != nil {
			return nil, status.Error(codes.Internal, "error checking unique constraints")
		}
		defer cursor.Close(context.TODO())

		// Check if any other user has the same email or username
		for cursor.Next(context.TODO()) {
			var otherUser types.User
			if err := cursor.Decode(&otherUser); err == nil {
				if email != nil && otherUser.Email == *email {
					return nil, status.Errorf(codes.AlreadyExists, "email %s already in use", *email)
				}
				if username != nil && otherUser.Username == *username {
					return nil, status.Errorf(codes.AlreadyExists, "username %s already in use", *username)
				}
			}
		}
	}

	update := bson.M{"$set": bson.M{}}
	if email != nil {
		update["$set"].(bson.M)["email"] = *email
	}
	if name != nil {
		update["$set"].(bson.M)["name"] = *name
	}
	if username != nil {
		update["$set"].(bson.M)["username"] = *username
	}
	if image != nil {
		update["$set"].(bson.M)["image"] = *image
	}

	var updatedUser types.User
	err = s.collection.FindOneAndUpdate(
		context.TODO(),
		bson.M{"id": id},
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&updatedUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating user %s", id)
	}

	var deletedAtTimestamp *timestamppb.Timestamp
	if updatedUser.DeletedAt != nil {
		deletedAtTimestamp = timestamppb.New(*updatedUser.DeletedAt)
	}

	return &pb.User{
		Id:             updatedUser.Id,
		Email:          updatedUser.Email,
		Name:           updatedUser.Name,
		Username:       updatedUser.Username,
		Image:          updatedUser.Image,
		PostCount:      updatedUser.PostCount,
		FollowerCount:  updatedUser.FollowerCount,
		FollowingCount: updatedUser.FollowingCount,
		CreatedAt:      timestamppb.New(updatedUser.CreatedAt),
		DeletedAt:      deletedAtTimestamp,
	}, nil
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
