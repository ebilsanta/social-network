package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ebilsanta/social-network/backend/user-service/api"
	pb "github.com/ebilsanta/social-network/backend/user-service/proto/generated"
	"github.com/ebilsanta/social-network/backend/user-service/storage"
	"google.golang.org/grpc"
)

func main() {
	store, err := storage.NewMongoStore()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := store.Client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Default().Println("User service running on port:", os.Getenv("SERVER_PORT"))

	followerClient, followerConn := api.InitFollowerService()
	defer func() {
		if err := followerConn.Close(); err != nil {
			log.Printf("Error closing FollowerService connection: %v", err)
		}
	}()

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, api.NewServer(store, followerClient))
	grpcServer.Serve(lis)
}
