// Code listing 16: https://play.golang.org/p/SKEOuo5WGD

package main

import "fmt"

func main() {
   // Array with 5 rows and 2 columns
   var arr = [5][2]int{ {0,0}, {2,4}, {1,3}, {5,7}, {6,8}}

   // Display each array element's value
   for i := 0; i < 5; i++ {
      for j := 0; j < 2; j++ {
         fmt.Printf("arr[%d][%d] = %d\n", i,j, arr[i][j] )
      }
   }
}