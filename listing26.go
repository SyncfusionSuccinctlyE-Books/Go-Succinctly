// Code listing 26: https://play.golang.org/p/NnwWpJzGVg

package main

import "fmt"

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

	for i := 1; i < 4; i++ {
		fmt.Printf("\nRUN NUMBER %d\n", i)
		for key, value := range actor {
			fmt.Printf("%s : %d years old\n", key, value)
		}
	}
}