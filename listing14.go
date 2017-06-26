// Code listing: https://play.golang.org/p/Fgx_WlXzzC

package main

import (
	"fmt"
	"os"
)

func main() {

	numAtoI, numJtoR, numStoZ, numSpaces, numOther := 0, 0, 0, 0, 0
	sentence := os.Args[1]

	for i := 1; i < len(sentence); i++ {
		letter := string(sentence[i])
		switch letter {
		case "a", "b", "c", "d", "e", "f", "g", "h", "i":
			numAtoI += 1
		case "j", "k", "l", "m", "n", "o", "p", "q", "r":
			numJtoR += 1
		case "s", "t", "u", "v", "w", "x", "y", "z":
			numStoZ += 1
		case " ":
			numSpaces += 1
		default:
			numOther += 1
		}
	}

	fmt.Printf("A to I: %d\n", numAtoI)
	fmt.Printf("J to R: %d\n", numJtoR)
	fmt.Printf("S to Z: %d\n", numStoZ)
	fmt.Printf("Spaces: %d\n", numSpaces)
	fmt.Printf("Other: %d\n", numOther)
}
