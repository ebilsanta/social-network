package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	services "github.com/ebilsanta/social-network/backend/complex_services/follower_complex_service/services"
)

func Init() {
	followerClient, followerConn := services.InitFollowerService()
	defer followerConn.Close()

	userClient, userConn := services.InitUserService()
	defer userConn.Close()

	quit := make(chan struct{})
	producer := services.StartKafkaProducer(os.Getenv("KAFKA_BROKER"), quit)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigchan
		close(quit)
		log.Println("Shutting down follower complex service producer")
	}()

	r := NewRouter(followerClient, userClient, producer)
	r.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
