// Code listing 50: https://play.golang.org/p/8iRk4YT9iw

package main

import (
	"io/ioutil"
	"os"
)

func main() {

	// Write a new file from a byte string
	name := "test.txt"
	txt := []byte("Not much in this file.")
	if err := ioutil.WriteFile(name, txt, 0755); err != nil {
		panic(err)
	}

	// Read the contents of a file into a []byte
	results, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	println(string(results))

	// Or use os.Open(filename)
	reader, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	results, err = ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	reader.Close()
	println(string(results))
}