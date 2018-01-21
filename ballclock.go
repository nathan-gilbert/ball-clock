package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//ClockState -- the current state of the ball clock
type ClockState struct {
	Min     []int //holds 4 balls
	FiveMin []int //holds 11 balls
	Hour    []int //holds 11 balls, but always contains 1 (so count starts @ 1)
	Main    []int //holds at most 'nBalls'
}

//timer function to use with defer
func timeIt(start time.Time) {
	elapsed := time.Since(start)
	ms := int64(elapsed) / int64(time.Millisecond)
	s := float64(elapsed) / float64(time.Nanosecond)
	fmt.Printf("Completed in %d milliseconds (%.3f seconds)", ms, s)
}

//TODO increments the clock state by 1 minute
func incrementState(state ClockState) ClockState {
	var newState ClockState
	return newState
}

//TODO implement mode1
func mode1(nBalls int) {
	defer timeIt(time.Now())
	days := 0
	state := initClockState(nBalls)
	//the first state
	initialState := state
	state = incrementState(state)
	//while the current state is not the same as the initial state
	//keep interating
	minutes := 0
	for !reflect.DeepEqual(initialState, state) { //reflection can be inefficent, write own
		state = incrementState(state)
		minutes++
	}
	days = minutes / (60 * 24)
	fmt.Printf("%d balls cycle after %d days.", nBalls, days)
}

//TODO implement mode2
func mode2(nBalls int, nMinutes int) {
	defer timeIt(time.Now())
	state := initClockState(nBalls)

	//simulate for nMinutes...
	for i := 0; i < nMinutes; i++ {
		state = incrementState(state)
	}

	//get the results
	jsonResult, err := json.Marshal(state)
	if err != nil {
		log.Fatal("unable to convert ClockState to JSON")
	}
	fmt.Printf("%s\n", jsonResult)
}

func initClockState(nBalls int) ClockState {
	var state ClockState
	//setup the main queue
	for i := 1; i <= nBalls; i++ {
		state.Main = append(state.Main, i)
	}
	return state
}

//check the cmdline args and modes
func parseArgs(options []string) (int, int) {
	if len(options) < 1 {
		log.Fatal("Not enough arguments.")
	}

	nBalls, err := strconv.Atoi(options[1])
	if err != nil {
		log.Fatal("Error converting nBalls to integer.")
	}

	//from the implementation specifications...
	if nBalls < 27 || nBalls > 127 {
		log.Fatal("Number of balls must be between 27 and 127 (inclusive)")
	}

	nMinutes := 0
	if strings.Compare(options[0], "-mode2") == 0 {
		nMinutes, err = strconv.Atoi(options[2])
		if err != nil {
			log.Fatal("Error converting nMinutes to integer.")
		}
	}
	return nBalls, nMinutes
}

func main() {
	/*
		arguments:
		-mode1 X
		-mode2 X [Y]
	*/
	options := os.Args[1:]
	nBalls, nMinutes := parseArgs(options)

	if strings.Compare(options[0], "-mode1") == 0 {
		mode1(nBalls)
	} else if strings.Compare(options[0], "-mode2") == 0 {
		mode2(nBalls, nMinutes)
	} else {
		fmt.Println("Unknown option. Must be -mode1 or -mode2")
		os.Exit(1)
	}
}
