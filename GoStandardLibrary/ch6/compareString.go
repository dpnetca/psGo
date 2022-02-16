package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "This is a string"
	// string2 := "This is another string"
	// string2 := "This is a string"
	string2 := "this is a string"

	if string1 == string2 { // case sensitive, comaresd byte values
		fmt.Println("identical")
	} else {
		fmt.Println("not identical")
	}

	// 0 = match,
	if strings.Compare(string1, string2) == 0 {
		fmt.Println("identical")
	} else {
		fmt.Println("not identical")
	}

	if CompareCaseIns(string1, string2) {
		fmt.Println("match")
	} else {
		fmt.Println("not match")
	}
}

func CompareCaseIns(a, b string) bool {
	if len(a) == len(b) {
		if strings.ToLower(a) == strings.ToLower(b) {
			return true
		}

	}
	return false
}
