package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "Some long string to do some thing fun"

	string2 := strings.ToLower(string1)
	fmt.Println(string2)

	string2 = strings.ToUpper(string1)
	fmt.Println(string2)

	string2 = strings.Title(string1)
	fmt.Println(string2)

	fmt.Println(properTitle(string1))

}

func properTitle(input string) string {
	words := strings.Fields(input)
	smallwords := " a an on the to "
	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}
