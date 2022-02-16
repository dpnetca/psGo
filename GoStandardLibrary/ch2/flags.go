package main

import (
	"flag"
	"fmt"
)

func main() {
	archPtr := flag.String("arch", "x86", "CPU Type")

	flag.Parse()

	switch *archPtr {
	case "x86":
		fmt.Println("Running 32 bit mode")
	case "AMD64":
		fmt.Println("Running 64 bit mode")
	case "IA64":
		fmt.Println("dude...")

	}
	fmt.Println("Process complete")
}
