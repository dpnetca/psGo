package main

import (
	"fmt"
	"regexp"
)

func main() {
	string1 := "some long string to do something fun"

	// r, _ := regexp.Compile(`s([a-z]+)e`)
	r, _ := regexp.Compile(`s(\w[a-z]+)e\b`)

	fmt.Println(r.MatchString(string1))
	fmt.Println(r.FindAllString(string1, -1))
	fmt.Println(r.FindAllStringIndex(string1, -1))

	newText := r.ReplaceAllString(string1, "a")
	fmt.Println(newText)
}
