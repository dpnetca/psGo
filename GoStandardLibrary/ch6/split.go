package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "This is a string"

	splitString1 := strings.Split(string1, " ")

	for i := range splitString1 {
		fmt.Printf("%v \n", splitString1[i])
	}
	fmt.Println()

	string2 := "This is a string|and another|one more time"

	splitString2 := strings.SplitAfter(string2, "|")

	for i := range splitString2 {
		fmt.Printf("%v \n", splitString2[i])
	}
	fmt.Println()

	splitString3 := strings.SplitAfterN(string1, " ", 2)

	for i := range splitString3 {
		fmt.Printf("%v \n", splitString3[i])
	}
	fmt.Println()
}
