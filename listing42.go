// Code listing 42: https://play.golang.org/p/8d9Vw9eiCH

package main

import (
	"fmt"
)

func generateAccountNumber(accountNumberChannel chan int) {
	// internal variable to store last generated account number
	var accountNumber int
	accountNumber = 30000001

	for {
		accountNumberChannel <- accountNumber
		// increment the account number by 1
		accountNumber += 1
	}
}

func main() {
	accountNumberChannel := make(chan int)
	// start the goroutine that generates account numbers
	go generateAccountNumber(accountNumberChannel)

	fmt.Printf("SMITH: %d\n", <-accountNumberChannel)
	fmt.Printf("SINGH: %d\n", <-accountNumberChannel)
	fmt.Printf("JONES: %d\n", <-accountNumberChannel)
	fmt.Printf("LOPEZ: %d\n", <-accountNumberChannel)
	fmt.Printf("CLARK: %d\n", <-accountNumberChannel)
}