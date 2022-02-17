package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func main() {
	// fmt.Println("Our current version of Go is " + runtime.Version()) // print line w/ newline at end
	// fmt.Print("Our current version of Go is " + runtime.Version()) // no newline

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	text, _ := reader.ReadString('\n') // note the return character becomes part of the string

	fmt.Printf("Hello %v", text)

	fmt.Printf("We are running GO %v in %v!\n", runtime.Version(), runtime.GOOS) // tokenize

}
