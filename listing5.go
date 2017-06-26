// Code listing 5: https://play.golang.org/p/aj8oXdf9_m

package main

import (
	"fmt"
)

func main() {
	msg := "Today's prize-winning entry is %d\n"
	winner := 9
	fmt.Printf(msg, winner)
}
