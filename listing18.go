// Code listing 18: https://play.golang.org/p/iWEMG5qjaG

package main

import (
	"fmt"
)

func main() {
	fruits := [...]string{"apples", "oranges", "bananas", "kiwis"}
	fmt.Printf("%v\n", fruits[1:3])
	fmt.Printf("%v\n", fruits[0:2]) 
	fmt.Printf("%v\n", fruits[:3]) 
	fmt.Printf("%v\n", fruits[2:]) 
}