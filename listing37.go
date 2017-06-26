// Code listing 37: https://play.golang.org/p/6TXx5SbqsU

package main

import (
	"fmt"
)

func main() {
	var anything interface{} = "something"
	aInt := anything.(int)
	fmt.Println(aInt)
}