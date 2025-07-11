package poker

import "sync"

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}, sync.RWMutex{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
	lock  sync.RWMutex
}

// assert InMemoryPlayerStore implements PlayerStore.
var _ PlayerStore = &InMemoryPlayerStore{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetLeague() League {
	var league League

	i.lock.RLock()
	defer i.lock.RUnlock()
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}
