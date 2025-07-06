package main

import (
	"io"
	"log"
	"os"
)

type tape struct {
	file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
	t.file.Truncate(0)
	if _, err := t.file.Seek(0, io.SeekStart); err != nil {
		log.Fatalf("could not seek file to start, %v", err)
	}
	return t.file.Write(p)
}
