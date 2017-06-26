// Code listing 35: https://play.golang.org/p/ghLsH47659

package main

import (
	"fmt"
)

func main() {
	displayType(42)
	displayType(3.14)
	displayType("ここでは文字列です")
}

func displayType(i interface{}) {
	switch theType := i.(type) {
	case int:
		fmt.Printf("%d is an §§§§§§§§§§§integer\n", theType)
	case float64:
		fmt.Printf("%f is a 64-bit float\n", theType)
	case string:
		fmt.Printf("%s is a string\n", theType)
	default:
		fmt.Printf("I don't know what %v is\n", theType)
	}
}