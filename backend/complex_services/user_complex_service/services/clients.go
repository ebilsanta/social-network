package api

import (
	"fmt"
	"log"
	"os"

	proto "github.com/ebilsanta/social-network/backend/complex_services/user_service/services/proto/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitUserService() (proto.UserServiceClient, *grpc.ClientConn) {
	userServiceAddr := fmt.Sprintf("user_service:%s", os.Getenv("USER_SVC_PORT"))

	conn, err := grpc.NewClient(userServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}
	userClient := proto.NewUserServiceClient(conn)
	return userClient, conn
}
