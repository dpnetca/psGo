package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "      some text   "
	fmt.Printf("%q\n", string1)
	string2 := strings.TrimSpace(string1)
	fmt.Printf("%q\n", string2)

	string2 = strings.TrimLeft(string1, " ")
	fmt.Printf("%q\n", string2)
	// TrimPrefix and TrimSuffix are useful as well...

}
