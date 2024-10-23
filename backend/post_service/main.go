package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ebilsanta/social-network/backend/post-service/api"
	"github.com/ebilsanta/social-network/backend/post-service/storage"
)

func main() {
	store, err := storage.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	quit := make(chan struct{})

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	producer := api.StartKafkaProducer(os.Getenv("KAFKA_BROKER"), "posts", quit)
	go api.StartGRPCServer(os.Getenv("SERVER_PORT"), store, producer, quit)

	<-sigchan
	close(quit)
	log.Println("Shutting down...")
}
