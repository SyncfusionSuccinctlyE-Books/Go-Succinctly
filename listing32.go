// Code listing 32: https://play.golang.org/p/HoT-yIfZpL

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
	
	fmt.Printf("\nDouble it!\n\n")
	r.double()
	fmt.Printf("Height: %d\n", r.height)
	fmt.Printf("Width: %d\n", r.width)
	fmt.Printf("Area: %d\n", r.area())
	
}

func (r rect) area() int {
	h := r.height
	w := r.width
	return h * w
}

func (r *rect) double() {
	r.height *= 2
	r.width *=2
}