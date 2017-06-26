// Code listing 20: https://play.golang.org/p/nETmIwxnOR

package main

import (
	"fmt"
)

func main() {
	mySlice := make([]int, 4, 8)
	fmt.Printf("Initial Length: %d\n", len(mySlice))
	fmt.Printf("Capacity: %d\n", cap(mySlice))
	fmt.Printf("Contents: %v\n", mySlice)
}