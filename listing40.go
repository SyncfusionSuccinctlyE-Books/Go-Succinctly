// Code listing 40: https://play.golang.org/p/8A-baEjZ_L

package main

import (
	"fmt"
)

func message(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {
	go message("goroutine")
	message("normal")
}