// Code listing 6: https://play.golang.org/p/27AfacoyoM

package main

import (
	"fmt"
)

const (
	prizeDay = "Wednesday"
	prizeFund = 10000
)

func main() {
	msg := "%s's prize-winning entry is %d and wins %d!!!\n"
	winner := 9
	fmt.Printf(msg, prizeDay, winner, prizeFund)
}
