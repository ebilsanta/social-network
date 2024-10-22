package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/ebilsanta/social-network/backend/user-service/proto"
	"google.golang.org/grpc"
)

func main() {
	store, err := NewMongoStore()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := store.client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Default().Println("User service running on port:", os.Getenv("SERVER_PORT"))
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, newServer(store))
	grpcServer.Serve(lis)
}
