package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	newPerson := Person{"Joe", "Smith", 33}

	// print values of struct
	fmt.Printf("%v\n", newPerson)

	// print var type
	fmt.Printf("%T\n", newPerson)

	// print binary of the int
	fmt.Printf("%b\n", 1234)

	// print hex of the int
	fmt.Printf("%x\n", 1234)

	// print character with ascii value 42 (*)
	fmt.Printf("%c\n", 42)

	// https://pkg.go.dev/fmt has a list of formatter options
}
