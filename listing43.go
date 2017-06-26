// Code listing 43: https://play.golang.org/p/xisjdH4Jn4

package main

import (
	"fmt"
)

func generateAccountNumber(accountNumberChannel chan int) {
	// internal variable to store last generated account number
	var accountNumber int
	accountNumber = 30000001

	for {
		// close the channel after 5 numbers are requested
		if accountNumber > 30000005 {
			close(accountNumberChannel)
			return
		}
		accountNumberChannel <- accountNumber
		// increment the account number by 1
		accountNumber += 1
	}
}

func main() {
	accountNumberChannel := make(chan int)
	// start the goroutine that generates account numbers
	go generateAccountNumber(accountNumberChannel)

	// slice containing new customer names
	newCustomers := []string{"SMITH", "SINGH", "JONES", "LOPEZ",
                                                 "CLARK", "ALLEN"}

	// get a new account number for each customer
	for _, newCustomer := range newCustomers {
		// is there anything to retrieve from the channel?
		accnum, ok := <-accountNumberChannel
		if !ok {
			fmt.Printf("%s: No number available\n",
                                                      newCustomer)
		} else {
			fmt.Printf("%s: %d\n", newCustomer, accnum)
		}
	}
}