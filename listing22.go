// Code listing 22: https://play.golang.org/p/u7HUmN6DQz

package main

import (
	"fmt"
)

func main() {

	mySlice := make([]int, 0, 8)
	mySlice = append(mySlice, 1, 3, 5, 7, 9, 11, 13, 17)

	fmt.Printf("Contents: %v\n", mySlice)
	fmt.Printf("Number of Items: %d\n", len(mySlice))
	fmt.Printf("Capacity: %d\n", cap(mySlice))

	mySlice[8] = 19

	fmt.Printf("Contents: %v\n", mySlice)
	fmt.Printf("Number of Items: %d\n", len(mySlice))
	fmt.Printf("Capacity: %d\n", cap(mySlice))

}