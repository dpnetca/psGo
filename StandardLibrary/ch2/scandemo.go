package main

import "fmt"

func main() {
	var name string
	fmt.Println("What is your name?")

	/*
		entering in multiple words causes....issues
		only read up to the space....
	*/
	// inputs, _ := fmt.Scanf("%s", &name)

	// %q looks for quoted string..weird things happen when no quotes entered though
	inputs, _ := fmt.Scanf("%q", &name)

	//bufio may be better fpor this

	fmt.Printf("inputs = %d \n", inputs)
	fmt.Printf("Hello %s\n", name)
}
