package main

import (
	"fmt"
	"session03/kvasync/storage"
)

const storeName = "store.db"

type entry struct {
	Key, Value string
}

type Store struct {
	storage storage.Storage
	allChan chan bool
	entryChan chan *entry
	reqChan chan string
	closeChan chan bool
}

func NewStore(storage storage.Storage) *Store {
	store := &Store{
		storage: storage,
		allChan: make(chan bool),
		entryChan: make(chan *entry),
		reqChan: make(chan string),
		closeChan: make(chan bool),
	}

	go store.run()

	return store
}

func NewDefaultStore() *Store {
	return NewStore(storage.NewFileStorage(storeName))
}

func (s *Store) Add(key, value string) {
	newEntry := &entry{
		Key: key,
		Value: value,
	}

	s.entryChan <- newEntry
}

func (s *Store) Get(key string) {
	s.reqChan <- key
}

func (s *Store) All() {
	s.allChan <- true
}

func (s *Store) run() {
	var newEntry *entry
	for {
		select {
		case newEntry = <-s.entryChan:
			err := s.storage.Write(newEntry.Key, newEntry.Value)
			if err != nil {
				fmt.Println(err)
			}
		case key := <-s.reqChan:
			keyValueMap, err := s.storage.Read()
			if err != nil {
				fmt.Println(err)
			}

			value, found := keyValueMap[key]
			if found {
				fmt.Println("> " + key + "=" + value)
			}
		case <-s.allChan:
			keyValueMap, err := s.storage.Read()
			if err != nil {
				fmt.Println(err)
			}

			for key := range keyValueMap {
				fmt.Println("> " + key + "=" + keyValueMap[key])
			}
		case <-s.closeChan:
			return
		}
	}
}

func (s *Store) Close() {
	s.closeChan <- true
}