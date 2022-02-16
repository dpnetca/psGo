package main

import "fmt"

func main() {
	age := 42
	// out, _ := fmt.Print("Jeremy is ", age, " years old \n")

	// print("Bytes written: ", out)

	name := "Jeremy"

	// %v works but specific is better
	// fmt.Printf("Teachers name is %v is %v years old \n", name, age)
	fmt.Printf("Teachers name is %s is %d years old \n", name, age)
	// age := "apple"
	// fmt.Printf("Teachers name is %s is %d years old \n", name, age)

	/*
		https://pkg.go.dev/fmt
		%v formats the valuye in default format
		%s string
		%s decimal integers
		%f or %g float
		%b base 2
		%o base 8
		%t boolean

	*/

	var pi float32 = 3.14592
	fmt.Printf("Pi is %.2f\n", pi)

	fmt.Printf("|%7.2f|%7.2f|%7.2f|\n", 23.1, 42.2, 12.33)
	fmt.Printf("|%-7.2f|%-7.2f|%-7.2f|\n", 2.93, 4.122, 121.2)
	// %<total digits>.<digits after decimal>f
	// %-<total digits>.<digits after decimal>f  <--- left justify
	fmt.Printf("|%7s|%7s|%7s|\n", "aaa", "bbb", "ccc")
	fmt.Printf("|%-7q|%-7q|%-7q|\n", "aaa", "bbb", "ccc")

	output := fmt.Sprintf("|%-7s|%-7s|%-7s|\n", "ddd", "eee", "ggg")
	fmt.Print(output)
}
