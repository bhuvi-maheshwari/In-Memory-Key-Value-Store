package main

import (
	"fmt"

	"github.com/bhuvi-maheshwari/in-memory-key-value-store/service"
)

func main() {

	store := service.NewStore()
	pairs := []service.Pair{
		{
			Key:   "country",
			Value: "India",
		},
		{
			Key:   "state",
			Value: "Karnataka",
		},
	}
	store.Put("Test_Key", pairs)

	pairs = []service.Pair{
		{Key: "title", Value: "SDE Bootcamp"},
		{Key: "price", Value: 3000.000},
		{Key: "enrolled", Value: false},
		{Key: "estimated_time", Value: 30},
	}
	store.Put("sde_bootcamp", pairs)
	fmt.Printf(store.Get("sde_bootcamp").String())
	fmt.Println("Keys: ", store.Keys())
	fmt.Println("Search: ", store.Search("country", "India"))
	fmt.Println("Delete: ", store.Delete("Test_Key"))

}
