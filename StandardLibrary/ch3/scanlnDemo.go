package main

import "fmt"

func main() {
	fmt.Print("What is your name: ")
	// var name string
	// fmt.Scanln(&name)
	// fmt.Printf("Hello %s nice to meet you!\n", name)

	var firstName string
	var lastName string
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hello %s %s nice to meet you!\n", firstName, lastName)
}
