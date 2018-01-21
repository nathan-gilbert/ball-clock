package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

//ClockState -- the current state of the ball clock
type ClockState struct {
	Min     []int
	FiveMin []int
	Hour    []int
	Main    []int
}

//TODO palante increments the clock state by 1 minute
func palante(state ClockState) ClockState {
	var newState ClockState
	return newState
}

//TODO implement mode1
func mode1(nBalls int) {
	//state := initClockState(nBalls)
}

//TODO implement mode2
func mode2(nBalls int, nMinutes int) ClockState {
	state := initClockState(nBalls)
	//setup the main queue
	for i := 1; i <= nBalls; i++ {
		state.Main = append(state.Main, i)
	}
	return state
}

func initClockState(nBalls int) ClockState {
	var state ClockState
	return state
}

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

	//check the cmdline args and modes
	nBalls, err := strconv.Atoi(options[1])
	if err != nil {
		log.Fatal("Error converting nBalls to integer.")
	}
	//from the implementation specifications...
	if nBalls < 27 || nBalls > 127 {
		log.Fatal("Number of balls must be between 27 and 127 (inclusive)")
	}

	ms := 0
	s := 0.0
	if strings.Compare(options[0], "-mode1") == 0 {
		start := time.Now()
		mode1(nBalls)
		elapsed := time.Since(start)
		ms = int(elapsed / time.Millisecond)
		s = float64(elapsed / time.Nanosecond)
	} else if strings.Compare(options[0], "-mode2") == 0 {
		nMinutes, err := strconv.Atoi(options[2])
		if err != nil {
			log.Fatal("Error converting nMinutes to integer.")
		}
		start := time.Now()
		finalState := mode2(nBalls, nMinutes)
		elapsed := time.Since(start)
		ms = int(elapsed / time.Millisecond)
		s = float64(elapsed / time.Nanosecond)
		jsonResult, err := json.Marshal(finalState)
		if err != nil {
			log.Fatal("unable to convert ClockState to JSON")
		}
		fmt.Printf("%s\n", jsonResult)
	} else {
		fmt.Println("Unknown option. Must be -mode1 or -mode2")
		os.Exit(1)
	}
	fmt.Printf("Completed in %d milliseconds (%2.3f seconds)", ms, s)
}
