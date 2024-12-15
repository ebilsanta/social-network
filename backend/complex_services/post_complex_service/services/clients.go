package services

import (
	"fmt"
	"log"
	"os"

	proto "github.com/ebilsanta/social-network/backend/complex_services/post_service/services/proto/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitPostService() (proto.PostServiceClient, *grpc.ClientConn) {
	postServiceAddr := fmt.Sprintf("post_service:%s", os.Getenv("POST_SVC_PORT"))

	conn, err := grpc.NewClient(postServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to post service: %v", err)
	}
	postClient := proto.NewPostServiceClient(conn)
	return postClient, conn
}

func InitUserService() (proto.UserServiceClient, *grpc.ClientConn) {
	userServiceAddr := fmt.Sprintf("user_service:%s", os.Getenv("USER_SVC_PORT"))

	conn, err := grpc.NewClient(userServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}
	userClient := proto.NewUserServiceClient(conn)
	return userClient, conn
}
