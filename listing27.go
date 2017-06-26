// Code listing 27: https://play.golang.org/p/k940jMZE4G

package main

import (
	"fmt"
	"sort"
)

func main() {
	actor := map[string]int{
		"Paltrow": 43,
		"Cruise":  53,
		"Redford": 79,
		"Diaz":    43,
		"Kilmer":  56,
		"Pacino":  75,
		"Ryder":   44,
	}

	// Store the keys in a slice
	var sortedActor []string
	for key := range actor {
		sortedActor = append(sortedActor, key)
	}
	// Sort the slice alphabetically
	sort.Strings(sortedActor)

	/* Retrieve the keys from the slice and use
	   them to look up the map values */
	for _, name := range sortedActor {
		fmt.Printf("%s : %d years old\n", name, actor[name])
	}
}