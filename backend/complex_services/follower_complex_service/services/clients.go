package services

import (
	"fmt"
	"log"
	"os"

	proto "github.com/ebilsanta/social-network/backend/complex_services/follower_complex_service/services/proto/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitFollowerService() (proto.FollowerServiceClient, *grpc.ClientConn) {
	followerServiceAddr := fmt.Sprintf("follower_service:%s", os.Getenv("FOLLOWER_SVC_PORT"))

	conn, err := grpc.NewClient(followerServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to follower service: %v", err)
	}
	followerClient := proto.NewFollowerServiceClient(conn)
	return followerClient, conn
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
