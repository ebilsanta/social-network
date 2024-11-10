package services

import (
	"fmt"
	"log"
	"os"

	proto "github.com/ebilsanta/social-network/backend/complex_services/feed_complex_service/services/proto/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitFeedService() (proto.FeedServiceClient, *grpc.ClientConn) {
	feedServiceAddr := fmt.Sprintf("feed_service:%s", os.Getenv("FEED_SVC_PORT"))

	conn, err := grpc.NewClient(feedServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to feed service: %v", err)
	}
	feedClient := proto.NewFeedServiceClient(conn)
	return feedClient, conn
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
