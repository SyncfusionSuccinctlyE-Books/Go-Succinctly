// Code listing 9: https://play.golang.org/p/k7ttmWQCA8

package main

import (
  "fmt"
  "reflect"
)

func main() {	
	fp1 := 22.0
	fp2 := 7.0
	
	result := fp1/fp2
	fmt.Printf("%f / %f = %f\n", fp1, fp2, result)
	  	  
  fmt.Printf("fp1 is type: %s\n", reflect.TypeOf(fp1))
  fmt.Printf("fp2 is type: %s\n", reflect.TypeOf(fp2))
  fmt.Printf("result is type: %s\n", reflect.TypeOf(result))
}
