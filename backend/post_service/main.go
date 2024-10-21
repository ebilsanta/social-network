package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/ebilsanta/social-network/backend/post-service/proto"
	"google.golang.org/grpc"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Default().Println("Post service running on port:", os.Getenv("SERVER_PORT"))
	grpcServer := grpc.NewServer()
	pb.RegisterPostServiceServer(grpcServer, newServer(store))
	grpcServer.Serve(lis)
}