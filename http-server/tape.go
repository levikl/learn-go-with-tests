package poker

import (
	"io"
	"log"
	"os"
)

type tape struct {
	file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
	if err := t.file.Truncate(0); err != nil {
		log.Fatalf("could not change size of file to 0, %v", err)
	}
	if _, err := t.file.Seek(0, io.SeekStart); err != nil {
		log.Fatalf("could not seek file to start, %v", err)
	}
	return t.file.Write(p)
}
