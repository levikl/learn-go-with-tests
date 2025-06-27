package main

import (
	"log"
	"os"

	blogposts "github.com/levikl/learn-go-with-tests/reading-files"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", posts)
}
