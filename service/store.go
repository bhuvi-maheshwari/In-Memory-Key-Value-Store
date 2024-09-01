package service

import (
	"sync"
)

type store struct {
	mu sync.Mutex
	m  map[string][]Pair
}

func NewStore() *store {
	return &store{
		m: make(map[string][]Pair),
	}
}

// Get the value of the key
func (s *store) Get(key string) Value {
	if pairs, ok := (*s).m[key]; ok {
		return pairs
	}
	return nil
}

// Put the key value pair
func (s *store) Put(key string, value Value) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

// Search and return all the keys with the given key and value
func (s *store) Search(key string, value interface{}) []string {
	keys := make([]string, 0)

	for k, v := range s.m {
		for _, val := range v {
			if val.Key == key && val.Value == value {
				keys = append(keys, k)
				break
			}
		}
	}
	return keys
}

// Get all the keys
func (s *store) Keys() []string {
	keys := make([]string, 0, len(s.m))
	for k := range s.m {
		keys = append(keys, k)
	}
	return keys
}

// Delete the key
func (s *store) Delete(key string) bool {
	if _, ok := (*s).m[key]; ok {
		delete(s.m, key)
		return true
	}
	return false
}
