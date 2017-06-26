// Code listing 44: https://play.golang.org/p/iO9J_0bW4r

package main

import (
	"fmt"
)

func main() {
	// un-buffered channel
	c := make(chan int)
	c <- 3
	fmt.Println("OK.")
}