// Code listing 34: https://play.golang.org/p/-g-IWBgLx4

package main

import (
	"fmt"
)

type Person struct {
	Name string
}

type Animal struct {
	Sound string
}

type Greeter interface {
	SayHi() string
}

// this method is specific to the Person class
func (p Person) SayHi() string {
	return "Hello! My name is " + p.Name
}

// this method is specific to the Animal class
func (a Animal) SayHi() string {
	return a.Sound
}

/* this method can be called on any type that
implements the Greeter interface */
func greet(i Greeter) {
	fmt.Println(i.SayHi())
}

func main() {
	man := Person{Name: "Bob Smith"}
	dog := Animal{Sound: "Woof! Woof!"}

	fmt.Println("\nPerson : ")
	greet(man)

	fmt.Println("\nAnimal : ")
	greet(dog)

}