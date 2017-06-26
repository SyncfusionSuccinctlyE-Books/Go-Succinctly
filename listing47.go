// Code listing 47: https://play.golang.org/p/X2q_LewMux

package main

import (
	"fmt"
	"time"
)

func broadcast(nsChannel chan int, cChannel chan bool) {

	numbers := []int{
		101,
		102,
		103,
		104,
		105,
		106,
		107,
		108,
		109,
		110,
	}
	i := 0
	for {
		// see which channel has items
		select {
		/* if the numbersChannel is being listened to,
		 take each number sequentially from the
		slice and put it into the channel */
		case nsChannel <- numbers[i]:
			i += 1
			/* if we've reached the last number and
			the channel is still being listened to,
			start reading from the beginning of the
			slice again */
			if i == len(numbers) {
				i = 0
			}
		/* if we receive a message on the
		complete channel, we stop transmitting */
		case <-cChannel:
			cChannel <- true
			return
		}
	}
}

func main() {
	numbersStation := make(chan int)
	completeChannel := make(chan bool)

	// execute broadcast in a separate thread
	go broadcast(numbersStation, completeChannel)

	// get 100 numbers from the numbersStation channel
	for i := 0; i < 100; i++ {
		// delay for artistic effect only
		time.Sleep(100 * time.Millisecond)
		// retrieve values from the channel
		fmt.Printf("%d ", <-numbersStation)
	}

	/* once we have received 100 numbers,
	send a message on completeChannel
	to tell it to stop transmitting */
	completeChannel <- true

	/* don't terminate the program until
	we receive a message on the completeChannel.
	Discard the response. */
	<-completeChannel

	/* we only get to here if we received a
	message on completeChannel */
	fmt.Println("Transmission Complete.")
}