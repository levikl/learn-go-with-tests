package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{}
	log.Fatal(http.ListenAndServe(":3210", server))
}
