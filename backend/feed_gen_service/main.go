package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ebilsanta/social-network/backend/feed-gen-service/api"
	"github.com/ebilsanta/social-network/backend/feed-gen-service/storage"
)

func main() {
	store, err := storage.NewRedisStore()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := store.Client.Close(); err != nil {
			log.Printf("Error closing Redis client: %v", err)
		}
	}()

	followerClient, followerConn := api.InitFollowerService()
	defer func() {
		if err := followerConn.Close(); err != nil {
			log.Printf("Error closing FollowerService connection: %v", err)
		}
	}()

	postClient, postConn := api.InitPostService()
	defer func() {
		if err := postConn.Close(); err != nil {
			log.Printf("Error closing PostService connection: %v", err)
		}
	}()

	quit := make(chan struct{})

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	consumer := api.StartKafkaConsumer(os.Getenv("KAFKA_BROKER"), quit)
	go api.StartGRPCServer(os.Getenv("SERVER_PORT"), store, followerClient, postClient, consumer, quit)

	<-sigchan
	close(quit)
	log.Println("Feed gen service shutting down...")
}
