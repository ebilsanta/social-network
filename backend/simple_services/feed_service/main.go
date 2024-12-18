package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ebilsanta/social-network/backend/feed-service/api"
	pb "github.com/ebilsanta/social-network/backend/feed-service/proto/generated"
	"github.com/ebilsanta/social-network/backend/feed-service/storage"
	"google.golang.org/grpc"
)

func main() {
	store, err := storage.NewRedisStore()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := store.Client.Close(); err != nil {
			log.Printf("Error closing Redis client: %v", err)
		}
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Default().Println("Feed service running on port:", os.Getenv("SERVER_PORT"))

	grpcServer := grpc.NewServer()
	pb.RegisterFeedServiceServer(grpcServer, api.NewServer(store))
	grpcServer.Serve(lis)
}
