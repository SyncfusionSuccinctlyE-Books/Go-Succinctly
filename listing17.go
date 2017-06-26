// Code listing 17: https://play.golang.org/p/iMCIH8oGMU

package main

import (
	"fmt"
)

func main() {
	myArray := [...]string{"Apples", "Oranges", "Bananas"}
	fmt.Printf("Initial array values: %v\n", myArray)
	myFunction(myArray)
	fmt.Printf("Final array values: %v\n", myArray)
}

func myFunction(arr [3]string) {
	// Change Oranges to Strawberries
	arr[1] = "Strawberries"
	fmt.Printf("Array values in myFunction(): %v\n", arr)
}