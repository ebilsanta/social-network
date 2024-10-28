package api

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	pb "github.com/ebilsanta/social-network/backend/user-service/proto/generated"
	"github.com/ebilsanta/social-network/backend/user-service/storage"
	"google.golang.org/grpc"
)

func StartGRPCServer(port string, store storage.Storage, followerClient pb.FollowerServiceClient, consumer *kafka.Consumer, quit chan struct{}) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	log.Default().Println("User service running on port:", os.Getenv("SERVER_PORT"))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	server := NewServer(store, followerClient, consumer)
	go server.StartUsersListener(quit)

	pb.RegisterUserServiceServer(grpcServer, server)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-quit
	grpcServer.Stop()
}
