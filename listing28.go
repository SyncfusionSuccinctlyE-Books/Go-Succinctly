// Code listing 28: https://play.golang.org/p/dJh0u0a8m6

package main

import "fmt"

func main() {
	a, b := 20, 71

	// pointer to a
	p := &a
	// read a via the pointer
	fmt.Printf("Value of a is: %d\n", *p)
	// set a via the pointer
	*p = 21
	// display the new value of a
	fmt.Printf("New value of a is: %d\n", a)

	// pointer to b
	p = &b
	// read b via the pointer
	fmt.Printf("Value of b is: %d\n", *p)
	// add 10 to b via its pointer
	*p = *p + 10
	// display the new value of b
	fmt.Printf("New value of b is: %d\n", b)
}

