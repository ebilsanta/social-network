package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/ebilsanta/social-network/backend/follower-service/proto/generated"
	"google.golang.org/grpc"
)

func main() {
	store, err := NewGraphStore()
	if err != nil {
		log.Fatal(err)
	}
	defer store.driver.Close(store.ctx)

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Default().Println("Follower service running on port:", os.Getenv("SERVER_PORT"))
	grpcServer := grpc.NewServer()
	pb.RegisterFollowerServiceServer(grpcServer, newServer(store))
	grpcServer.Serve(lis)
}
