package api

import (
	"fmt"
	"log"
	"os"

	proto "github.com/ebilsanta/social-network/backend/feed-gen-service/api/proto/generated"
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

func InitPostService() (proto.PostServiceClient, *grpc.ClientConn) {
	postServiceAddr := fmt.Sprintf("post_service:%s", os.Getenv("POST_SVC_PORT"))

	conn, err := grpc.NewClient(postServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to post service: %v", err)
	}
	postClient := proto.NewPostServiceClient(conn)
	return postClient, conn
}
