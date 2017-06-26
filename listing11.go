// Code listing 11: https://play.golang.org/p/vkpZgoCVk2

package main

import (
	"fmt"
	"os"
)

func main() {
    
  val1 := os.Args[1]
  val2 := os.Args[2]
  val3 := os.Args[3]
  val4 := os.Args[4]

  fmt.Printf("%s %s %s %s\n", val1, val2, val3, val4)
}


