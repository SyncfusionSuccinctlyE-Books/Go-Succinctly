// Code listing 45: https://play.golang.org/p/6wq2bE1Xch

package main

import (
	"fmt"
)

func main() {
	// buffered channel, capacity 1
	c := make(chan int, 1)
	c <- 3
	fmt.Println("OK.")
}