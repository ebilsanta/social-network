package api

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/ebilsanta/social-network/backend/follower-service/proto/generated"
	"github.com/ebilsanta/social-network/backend/follower-service/storage"
	"google.golang.org/grpc"
)

func StartGRPCServer(port string, store storage.Storage, producer *KafkaClient, quit chan struct{}) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	log.Default().Println("Follower service running on port:", os.Getenv("SERVER_PORT"))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(loggingInterceptor))
	server := NewServer(store, producer)
	go server.ListenKafkaEvents(quit)
	pb.RegisterFollowerServiceServer(grpcServer, server)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-quit
	grpcServer.Stop()
}

func loggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	// Log the incoming request
	log.Printf("Received request for method: %s", info.FullMethod)

	// Call the handler to execute the actual method
	resp, err := handler(ctx, req)

	// Log the response or error
	if err != nil {
		log.Printf("Error calling method %s: %v", info.FullMethod, err)
	} else {
		log.Printf("Successfully called method %s", info.FullMethod)
	}

	return resp, err
}
