package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/levikl/learn-go-with-tests/http-server"
)

const dbFileName = "game.db.json"

func main() {
	store, cleanup, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	poker.NewCLI(store, os.Stdin, &poker.SpyBlindAlerter{}).PlayPoker()
}
