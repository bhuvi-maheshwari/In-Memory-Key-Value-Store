package main

import (
	"fmt"
	"strings"
)

type Pair struct {
	key   string
	value interface{}
}

type Store map[string][]Pair
type Value []Pair

// Prints the value of the key
func (v Value) String() string {
	var sb strings.Builder
	for _, val := range v {
		s := fmt.Sprintf("%v: %v, ", val.key, val.value)
		sb.WriteString(s)
	}
	sb.WriteString("\n")
	return sb.String()
}

// Get the value of the key
func (s *Store) get(key string) Value {
	if pairs, ok := (*s)[key]; ok {
		return pairs
	}
	return nil
}

// Put the key value pair
func (s *Store) put(key string, value Value) {
	(*s)[key] = value
}

// Search and return all the keys with the given key and value
func (s *Store) search(key string, value interface{}) []string {
	keys := make([]string, 0)

	for k, v := range *s {
		for _, val := range v {
			if val.key == key && val.value == value {
				keys = append(keys, k)
				break
			}
		}
	}
	return keys
}

// Get all the keys
func (s *Store) keys() []string {
	keys := make([]string, 0, len(*s))
	for k := range *s {
		keys = append(keys, k)
	}
	return keys
}

// Delete the key
func (s *Store) delete(key string) bool {
	if _, ok := (*s)[key]; ok {
		delete(*s, key)
		return true
	}
	return false
}

func main() {
	pairs := []Pair{
		{
			key:   "country",
			value: "India",
		},
		{
			key:   "state",
			value: "Karnataka",
		},
	}

	store := Store{}
	store.put("Test_Key", pairs)

	pairs = []Pair{
		{key: "title", value: "SDE Bootcamp"},
		{key: "price", value: 3000.000},
		{key: "enrolled", value: false},
		{key: "estimated_time", value: 30},
	}
	store.put("sde_bootcamp", pairs)
	fmt.Printf(store.get("sde_bootcamp").String())
	fmt.Println("Keys: ", store.keys())
	fmt.Println("Search: ", store.search("country", "India"))
	fmt.Println("Delete: ", store.delete("Test_Key"))

}
