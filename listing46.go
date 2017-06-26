// Code listing 46: https://play.golang.org/p/Id7ulCJfsu

package main

import (
	"fmt"
)

func main() {
	// buffered channel, capacity 3
	c := make(chan int, 3)
	c <- 3
	c <- 4
	c <- 5
	c <- 6
	fmt.Println("OK.")
}