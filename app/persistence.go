package main

import "sync"

type Storage struct {
	mu sync.Mutex
	v  map[string]string
}

var (
	db = &Storage{v: make(map[string]string)}
)

func (s *Storage) Get(key string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.v[key]
}

func (s *Storage) Set(key string, value string) {
	s.mu.Lock()
	s.v[key] = value
	s.mu.Unlock()
}
