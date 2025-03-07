package main

import (
	"sync"
	"time"
)

type Entry struct {
	value  string
	ts     time.Time
	expiry int64
}
type Storage struct {
	mu sync.Mutex
	v  map[string]Entry
}

func (e Entry) expired() bool {
	return time.Now().Sub(e.ts).Milliseconds() > e.expiry
}

var (
	db = &Storage{v: make(map[string]Entry)}
)

func (s *Storage) Get(key string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	entry, present := s.v[key]
	if !present {
		return ""
	} else if entry.expiry == 0 {
		return entry.value
	} else if entry.expired() {
		return ""
	} else {
		return entry.value
	}
}

func (s *Storage) Set(key string, value string) {
	s.mu.Lock()
	s.v[key] = Entry{value: value, ts: time.Now()}
	s.mu.Unlock()
}

func (s *Storage) SetWithExpiry(key string, value string, expiry int64) {
	s.mu.Lock()
	s.v[key] = Entry{value: value, ts: time.Now(), expiry: expiry}
	s.mu.Unlock()
}
