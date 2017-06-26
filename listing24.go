// Code listing 24: https://play.golang.org/p/8ef9G_NIvN

package main

import "fmt"

func main() {
	actor := make(map[string]int)
	actor["Paltrow"] = 43
	actor["Cruise"] = 53
	actor["Redford"] = 79
	actor["Diaz"] = 43
	actor["Kilmer"] = 56
	actor["Pacino"] = 75
	actor["Ryder"] = 44

	fmt.Printf("Robert Redford is %d years old\n",
                                               actor["Redford"])
	fmt.Printf("Cameron Diaz is %d years old\n", actor["Diaz"])
	fmt.Printf("Val Kilmer is %d years old\n", actor["Kilmer"])
}