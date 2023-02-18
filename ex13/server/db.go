package server

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

type memoryDB struct {
	items map[string]string
	mu    *sync.RWMutex
}

func newDB() memoryDB {
	f, err := os.Open("db.json")
	if err != nil {
		return memoryDB{items: map[string]string{}}
	}

	items := map[string]string{}
	if err := json.NewDecoder(f).Decode(&items); err != nil {
		fmt.Println("could not decode data from persistent file storage to in-memory")
		return memoryDB{items: map[string]string{}}
	}
	return memoryDB{items: items}
}

func (m memoryDB) save() {
	f, err := os.Create("db.json")
	if err != nil {
		log.Fatalf("could not create the persistent file storage in fs\n", err.Error())
	}

	if err = json.NewEncoder(f).Encode(m.items); err != nil {
		log.Fatalf("could not encode date from in-memory to save in persistent file storage\n", err.Error())
	}
}

func (m memoryDB) set(key, value string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.items[key] = value
}

func (m memoryDB) get(key string) (string, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	value, found := m.items[key]
	return value, found
}

func (m memoryDB) delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.items, key)
}
