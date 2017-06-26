// Code listing 38: https://play.golang.org/p/PpEnpR1--y

package main

import (
	"fmt"
)

func main() {
	var anything interface{} = "something"
	aInt, ok := anything.(int)
	if !ok {
		fmt.Println("Cannot turn input into an integer")
	} else {
		fmt.Println(aInt)
	}
}