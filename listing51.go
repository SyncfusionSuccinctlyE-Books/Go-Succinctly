// Code listing 51: https://play.golang.org/p/QzJjwV9W6d

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
		// does "test" contain "es"?
		strings.Contains("test", "es"),

		// does "test" begin with "te"?
		strings.HasPrefix("test", "te"),

		// does "test" end in "st"?
		strings.HasSuffix("test", "st"),

		// how many times is "t" in test?
		strings.Count("test", "t"),

		// at what position is "e" in "test"?
		strings.Index("test", "e"),

		// join "input" and "output" with "/"
		strings.Join([]string{"input", "output"}, "/"),

		// repeat "Golly" 6 times
		strings.Repeat("Golly", 6),

		/* replace "xxxx" with the first two
		   non-overlapping instances of "a" replaced by "b" */
		strings.Replace("xxxx", "a", "b", 2),

		/* put "a-b-c-d-e" into a slice using
		   "-" as a delimiter */
		strings.Split("a-b-c-d-e", "-"),

		// put "TEST" in lower case
		strings.ToLower("TEST"),

		// put "TEST" in upper case
		strings.ToUpper("test"),
	)
}