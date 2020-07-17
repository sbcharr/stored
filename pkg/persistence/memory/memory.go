package memory

import (
	"errors"
	"sync"
)

// ErrKeyNotFound is called when no such key is available
var ErrKeyNotFound = errors.New("no such key found")

// Store is the main KV store
type Store struct {
	m     sync.RWMutex
	store map[string]string
}

// NewKVStore is the constructor
func NewKVStore() *Store {
	return &Store{store: make(map[string]string)}
}

// Put inserts a new record, else updates the existing record
func (s Store) Put(key, value string) error {
	s.m.Lock()
	defer s.m.Unlock()
	s.store[key] = value

	return nil
}

// Get retrieves an existing element
func (s Store) Get(key string) (string, error) {
	if value, ok := s.store[key]; ok {
		return value, nil
	}

	return "", ErrKeyNotFound
}

// Delete deletes an existing record
func (s Store) Delete(key string) error {
	if _, ok := s.store[key]; !ok {
		return ErrKeyNotFound
	}
	s.m.Lock()
	defer s.m.Unlock()
	delete(s.store, key)

	return nil
}
