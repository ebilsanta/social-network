package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ebilsanta/social-network/backend/follower-service/api"
	pb "github.com/ebilsanta/social-network/backend/follower-service/proto/generated"
	"github.com/ebilsanta/social-network/backend/follower-service/storage"
	"google.golang.org/grpc"
)

func main() {
	store, err := storage.NewGraphStore()
	if err != nil {
		log.Fatal(err)
	}
	defer store.Driver.Close(store.Ctx)

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Default().Println("Follower service running on port:", os.Getenv("SERVER_PORT"))
	grpcServer := grpc.NewServer()
	pb.RegisterFollowerServiceServer(grpcServer, api.NewServer(store))
	grpcServer.Serve(lis)
}
