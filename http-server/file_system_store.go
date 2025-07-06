package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"sync"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
	lock     sync.RWMutex
}

// assert FileSystemPlayerStore implements PlayerStore.
var _ PlayerStore = &FileSystemPlayerStore{}

func initializePlayerDBFile(file *os.File) error {
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		log.Fatalf("error calling file.Seek(0, io.SeekStart) in file %s, %v", file.Name(), err)
	}

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		if _, err := file.Write([]byte("[]")); err != nil {
			log.Fatalf("not able to write `[]` to file %s, %v", file.Name(), err)
		}
		if _, err := file.Seek(0, io.SeekStart); err != nil {
			log.Fatalf("error calling file.Seek(0, io.SeekStart) in file %s, %v", file.Name(), err)
		}
	}

	return nil
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initializePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem initializing player db file, %v", err)
	}

	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
		lock:     sync.RWMutex{},
	}, nil
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	f.lock.RLock()
	defer f.lock.RUnlock()

	if player := f.league.Find(name); player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	f.lock.Lock()
	defer f.lock.Unlock()

	if player := f.league.Find(name); player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	if err := f.database.Encode(f.league); err != nil {
		log.Fatalf("error encoding league into database, %v", err)
	}
}

func (f *FileSystemPlayerStore) GetLeague() League {
	f.lock.RLock()
	defer f.lock.RUnlock()

	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}
