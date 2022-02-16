package main

import "fmt"

func main() {
	ourString := "Go is Awesome"

	for i := 0; i < len(ourString); i++ {
		fmt.Printf("%x ", ourString[i])
	}
	fmt.Println()

	blogString := "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println(blogString)
	for i := 0; i < len(blogString); i++ {
		fmt.Printf("%q \n", blogString[i])
	}

	fmt.Println(ourString[1])
	fmt.Printf("%c\n", ourString[1])
	fmt.Println(ourString[0:5])
}
