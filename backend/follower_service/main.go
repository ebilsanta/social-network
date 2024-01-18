package main

import (
	"log"
	"os"
)

func main() {
	store, err := NewGraphStore()
	if err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":"+os.Getenv("SERVER_PORT"), store)
	server.Run()
}
