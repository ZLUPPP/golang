package main

import (
	"fmt"
	"strings"
)

func main() {
	m := map[string]int{"one": 1, "two": 2}

	// CRUD
	m["three"] = 3          // Create
	fmt.Println(m["three"]) // Read
	m["three"]++            // Update
	delete(m, "one")        // Delete

	fmt.Println(m)

	one, exist := m["two"]
	if exist {
		fmt.Println(one)
	}

	// Create an inventory management program for a small store.

	inventory := map[string]int{"apples": 10, "bananas": 5}
	fmt.Println(inventory["apples"])
	inventory["bananas"] = 12
	inventory["oranges"] = 8
	delete(inventory, "apples")

	for key, value := range inventory {
		fmt.Printf("%q: %d\n", key, value)
	}

	s := "Go for a walk"
	count := map[string]int{}
	for _, val := range strings.Split(s, " ") {
		count[val]++
	}

	fmt.Println(count)
}
