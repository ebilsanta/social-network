package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ebilsanta/social-network/backend/follower-service/api"
	"github.com/ebilsanta/social-network/backend/follower-service/storage"
)

func main() {
	store, err := storage.NewGraphStore()
	if err != nil {
		log.Fatal(err)
	}
	defer store.Driver.Close(store.Ctx)

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	quit := make(chan struct{})
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	kafkaClient := api.InitKafka(os.Getenv("KAFKA_BROKER"), quit)
	go api.StartGRPCServer(os.Getenv("SERVER_PORT"), store, kafkaClient, quit)

	<-sigchan
	close(quit)
	log.Println("Shutting down...")
}
