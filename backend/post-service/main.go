package main

import (
	"log"
	"os"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(os.Getenv("SERVER_PORT"), store)
	server.Run()
}
