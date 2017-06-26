// Code listing 15: https://play.golang.org/p/J7J4JQZSoY

package main

import (
	"fmt"
)

func main() {    

	total := 0
	mean := 0
	rainfallStats := [5]int{1091, 2010, 995, 1101, 1111}
	for _, value := range rainfallStats {
		total += value
	}
	mean = total / len(rainfallStats)
	fmt.Printf("Average rainfall: %d mm\n", mean)
}