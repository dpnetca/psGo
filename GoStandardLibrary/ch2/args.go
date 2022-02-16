package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 1 && args[0] == "/help" {
		fmt.Println("Usage: args <total meal> <tip percentage>")
		fmt.Println("Example: args 20 10")
	} else if len(args) != 2 {
		fmt.Println("You must enter 2 arguments type /help for help")
	} else {
		mealTotal, _ := strconv.ParseFloat(args[0], 32)
		tipAmount, _ := strconv.ParseFloat(args[1], 32)
		fmt.Printf("\nYour meal total will be %.2f\n", calculateTotal(float32(mealTotal), float32(tipAmount)))
	}
}

func calculateTotal(mealTotal float32, tipAmount float32) float32 {
	// using float to calculate currency is bad practice to to accuracy issues
	totalPrice := mealTotal + (mealTotal * tipAmount / 100)
	return totalPrice
}
