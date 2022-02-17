package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	fmt.Print(t, "\n")
	fmt.Print(t.Year(), "\n")
	fmt.Print(t.Month(), "\n")
	fmt.Print(t.Day(), "\n")

	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Format(time.UnixDate))
	fmt.Println(t.Format(time.Kitchen))

	// to make custom format use the following reference time:
	// Mon Jan 2 15:04:05 MST 2006
	fmt.Println(t.Format("Mon Jan 2 15:04:05 MST 2006"))
	fmt.Println(t.Format("MST Jan 2 2006 MON 15:04:05"))
	fmt.Println(t.Format("Monday, January 2 in the year 2006"))
	// easiest way would be assign format to a constant and referencing the constant instead of hard coding

	startDate := time.Date(2018, 07, 04, 9, 00, 00, 00, time.UTC)
	fmt.Println(startDate)

}
