package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

//classes
// balls
// the clock
// a way to advance time (probably a method on the clock)

func main() {
	/*
		arguments:
		-mode1 X
		-mode2 X [Y]
	*/
	options := os.Args[1:]
	if len(options) < 1 {
		fmt.Println("Not enough arguments.")
		os.Exit(0)
	}
	fmt.Println(options)

	if strings.Compare(options[0], "-mode1") == 0 {
		start := time.Now()
		//TODO implement mode1
		elapsed := time.Since(start)
		ms := int64(elapsed / time.Millisecond)
		s := float64(elapsed / time.Nanosecond)
		fmt.Printf("Completed in %d milliseconds (%2.3f seconds)", ms, s)
	} else if strings.Compare(options[0], "-mode2") == 0 {
		start := time.Now()
		//TODO implement mode2
		elapsed := time.Since(start)
		ms := int64(elapsed / time.Millisecond)
		s := float64(elapsed / time.Nanosecond)
		fmt.Printf("Completed in %d milliseconds (%2.3f seconds)", ms, s)
	} else {
		fmt.Println("Unknown option. Must be -mode1 or -mode2")
		os.Exit(1)
	}
}
