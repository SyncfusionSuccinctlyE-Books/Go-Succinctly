// Code listing 30: https://play.golang.org/p/ESEjDJcO5j

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
	fmt.Printf("Height: %d\n", r.height)
	fmt.Printf("Width: %d\n", r.width)
	fmt.Printf("Area: %d\n", r.area())
}

func (r rect) area() int {
	h := r.height
	w := r.width
	return h * w
}