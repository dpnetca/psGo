package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "some long string to do some thing fun"

	searchTerm := "long"
	result := strings.Contains(string1, searchTerm)
	fmt.Printf("result: %v\n", result)

	searchPrefix := "some long"
	result2 := strings.HasPrefix(string1, searchPrefix)
	fmt.Printf("result: %v\n", result2)

	string2 := strings.Replace(string1, "some", "a", -1)
	fmt.Printf("result: %v\n", string2)
}
