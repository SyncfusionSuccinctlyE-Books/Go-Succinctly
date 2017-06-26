// Code listing 7: https://play.golang.org/p/QBAB_gDZkW

package main

import (
	"fmt"
)

var (
	prizeDay = "Wednesday"
	prizeFund = 10000
)

func main() {
	prizeDay = "Thursday"
	prizeFund = 50000
	msg := "%s's prize-winning entry is %d and wins %d!!!\n"
	winner := 9
	fmt.Printf(msg, prizeDay, winner, prizeFund)
}
