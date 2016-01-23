package main

import (
	"session02/kvoop/storage"
)

const storeName = "store.db"

type Store struct {
	storage.Storage
}

func New(s storage.Storage) *Store {
	return &Store{s}
}

func NewFileStore(name string) (*Store, error) {
	fileStorage, err := storage.NewFileStorage(name)
	if err != nil {
		return nil, err
	}

	return New(fileStorage), nil
}

func NewDefaultStore() (*Store, error) {
	return NewFileStore(storeName)
}

func (s *Store) GetMultiple(keys []string) storage.KeyValueMap {
	subMap := storage.KeyValueMap{}
	for _, key := range keys {
		value, found := s.Get(key)
		if found {
			subMap[key] = value
		}
	}

	return subMap
}
