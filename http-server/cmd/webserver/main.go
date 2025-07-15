package main

import (
	"github.com/levikl/learn-go-with-tests/http-server"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	store, cleanup, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":3210", server); err != nil {
		log.Fatalf("could not listen on port 3210 %v", err)
	}
}
