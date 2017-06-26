// Code listing 13: https://play.golang.org/p/674IrwC_w1

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {    
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s is %s\n", os.Args[i],
                      oddOrEven(os.Args[i]))
	}
}

func oddOrEven(value string) string {
	num, _ := strconv.Atoi(value)
	if num % 2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}
