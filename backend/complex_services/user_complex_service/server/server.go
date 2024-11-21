package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ebilsanta/social-network/backend/complex_services/user_service/services"
)

func Init() {
	userClient, conn := services.InitUserService()
	defer conn.Close()

	followerClient, conn := services.InitFollowerService()
	defer conn.Close()

	quit := make(chan struct{})
	producer := services.StartKafkaProducer(os.Getenv("KAFKA_BROKER"), quit)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigchan
		close(quit)
		log.Println("Shutting down user complex service producer")
	}()

	r := NewRouter(userClient, followerClient, producer)

	r.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
