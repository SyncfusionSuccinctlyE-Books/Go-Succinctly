// Code listing 10: https://play.golang.org/p/S15y4hTPal

package main

import (
	"fmt"
	"strconv"
)

func main() {    
  int1 := 48
  int1_as_string := strconv.Itoa(int1)
  str1 := "My number is "
  str1 = str1 + int1_as_string
  fmt.Println(str1)

	str2 := "4"
  str2_as_integer, _ := strconv.Atoi(str2)
  result := 48 / str2_as_integer
  fmt.Printf("%s / %s is %d\n", int1_as_string, str2, result)	
}