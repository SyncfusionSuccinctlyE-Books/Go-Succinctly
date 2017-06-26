// Code listing 36: https://play.golang.org/p/JlO2GiWur4

package main

import (
	"fmt"
)

func main() {
	var anything interface{} = "something"
	aString := anything.(string)
	fmt.Println(aString)
}