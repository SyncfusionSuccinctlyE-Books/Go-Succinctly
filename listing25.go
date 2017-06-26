// Code listing 25: https://play.golang.org/p/oAcnEfSuBS

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

	fmt.Printf("Robert Redford is %d years old\n",
                                               actor["Redford"])
	fmt.Printf("Cameron Diaz is %d years old\n", actor["Diaz"])
	fmt.Printf("Val Kilmer is %d years old\n", actor["Kilmer"])
}
