package main

import (
	"fmt"
	"time"
)

func main() {
	// Formatting Reference Time:
	// Mon Jan 2 15:04:05 MST 2006

	first := time.Now()
	fmt.Printf("It is currently %v \n", first.Format("15:04:05"))
	// time.Sleep(2000000000)  // 2 seconds in nano seconds
	time.Sleep(1 * time.Second) // this is easier
	second := time.Now()
	fmt.Printf("It is now %v \n", second.Format("15:04:05"))

	startDate := time.Date(2018, 07, 04, 9, 00, 00, 00, time.UTC)

	elapsed := time.Since(startDate)

	fmt.Printf("%v has passed since %v\n", elapsed, startDate.Format("Monday Jan 2, 2006"))
	fmt.Printf("%v hours\n", elapsed.Hours())
	fmt.Printf("%.2f hours\n", elapsed.Hours())

	future := first.AddDate(0, 6, 0)
	fmt.Printf("in 6 months it will be %v\n", future)
	past := first.AddDate(0, -6, 0)
	fmt.Printf("6 months ago it was be %v\n", past)

	futureTime := first.Add(6 * time.Hour)
	fmt.Printf("in 6 hours it will be %v\n", futureTime.Format("15:04"))

	fmt.Printf("time until future time: %v\n", time.Until(futureTime))

}
