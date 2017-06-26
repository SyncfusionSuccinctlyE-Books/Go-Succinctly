// Code listing 23: https://play.golang.org/p/ukWiuzIO5h

package main

import (
	"fmt"
)

func main() {
	mySlice := make([]int, 0, 8)
	mySlice = append(mySlice, 1, 3, 5, 7, 9, 11, 13, 17)

	mySliceCopy := make([]int, 8)
	copy(mySliceCopy, mySlice)

	mySliceCopy[3] = 999
	fmt.Printf("mySlice: %v\n", mySlice)
	fmt.Printf("mySliceCopy: %v\n", mySliceCopy)
}