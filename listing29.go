// Code listing 29: https://play.golang.org/p/pAt9sdQYjP

package main

import (
	"fmt"
)

type rect struct {
	height int
	width  int
}

func main() {
	r := rect{height: 12, width: 20}
	fmt.Printf("Height: %d\nWidth: %d\n", r.height, r.width)
}