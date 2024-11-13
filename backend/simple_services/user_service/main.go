package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ebilsanta/social-network/backend/user-service/api"
	"github.com/ebilsanta/social-network/backend/user-service/storage"
)

func main() {
	store, err := storage.NewMongoStore()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := store.Client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan struct{})

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	consumer := api.StartKafkaConsumer(os.Getenv("KAFKA_BROKER"), quit)
	go api.StartGRPCServer(os.Getenv("SERVER_PORT"), store, consumer, quit)

	<-sigchan
	close(quit)
	log.Println("User service shutting down...")
}
