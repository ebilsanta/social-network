package api

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/ebilsanta/social-network/backend/post-service/proto"
	"github.com/ebilsanta/social-network/backend/post-service/storage"
	"google.golang.org/grpc"
)

func StartGRPCServer(port string, store storage.Storage, producer *KafkaProducer, quit chan struct{}) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	log.Default().Println("Post service running on port:", os.Getenv("SERVER_PORT"))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPostServiceServer(grpcServer, NewServer(store, producer))

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-quit
	grpcServer.Stop()
}
