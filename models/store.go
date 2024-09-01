package models

import "sync"

// type Store map[string][]Pair
type Store struct {
	Mu sync.Mutex
	M  map[string][]Pair
}
