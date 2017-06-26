// Code listing 52: https://play.golang.org/p/d0Ymnn_k85

package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("error message")
	fmt.Println(err)
}