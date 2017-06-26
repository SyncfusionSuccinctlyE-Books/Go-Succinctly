// Code listing 41: https://play.golang.org/p/lVR3_5JMfN

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func broadcast(c chan int) {
	// infinite loop to create random numbers
	for {
		/* generate a random number 0-999
		and put it into the channel */
		c <- rand.Intn(999)
	}
}

func main() {
	numbersStation := make(chan int)

	// execute broadcast in a separate thread
	go broadcast(numbersStation)

	// retrieve values from the channel
	for num := range numbersStation {
		// delay for artistic effect only
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("%d ", num)
	}
}