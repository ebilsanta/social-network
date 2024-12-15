package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	services "github.com/ebilsanta/social-network/backend/complex_services/post_service/services"
)

func Init() {
	postClient, postConn := services.InitPostService()
	defer postConn.Close()

	userClient, userConn := services.InitUserService()
	defer userConn.Close()

	quit := make(chan struct{})
	producer := services.StartKafkaProducer(os.Getenv("KAFKA_BROKER"), quit)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigchan
		close(quit)
		log.Println("Shutting down post complex service producer")
	}()

	r := NewRouter(postClient, userClient, producer)
	r.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
