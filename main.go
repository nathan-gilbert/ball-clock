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

//timer function to use with defer
func timeIt(start time.Time) {
	elapsed := time.Since(start)
	ms := int64(elapsed) / int64(time.Millisecond)
	s := float64(elapsed) / float64(time.Second)
	fmt.Printf("Completed in %d milliseconds (%.3f seconds)\n", ms, s)
}

func mode1(nBalls int) {
	defer timeIt(time.Now())
	days := 0
	var state ClockState
	var initialState ClockState
	state.Init(nBalls)
	initialState.Init(nBalls)

	//bootstrap by a minute to start off
	minutes := 1
	state.IncrementByMinute()
	//keep iterating while the current state is not the same as the initial state
	for !reflect.DeepEqual(initialState, state) { //reflection can be inefficient, maybe write own check?
		state.IncrementByMinute()
		minutes++
	}
	days = minutes / (60 * 24)
	fmt.Printf("%d balls cycle after %d days.\n", nBalls, days)
}

func mode2(nBalls int, nMinutes int) {
	defer timeIt(time.Now())
	var state ClockState
	state.Init(nBalls)

	//simulate for nMinutes...
	for i := 0; i < nMinutes; i++ {
		state.IncrementByMinute()
	}

	//get the results
	jsonResult, err := json.Marshal(state)
	if err != nil {
		log.Fatal("unable to convert ClockState to JSON")
	}
	fmt.Printf("%s\n", jsonResult)
}

//check the cmdline args and modes
func parseArgs(options []string) (int, int) {
	if len(options) < 1 {
		log.Fatal("Not enough arguments.")
	}

	var nBalls int
	var err error
	if len(options) > 1 {
		nBalls, err = strconv.Atoi(options[1])
		if err != nil {
			log.Fatal("Error converting nBalls to integer.")
		}
	} else {
		log.Fatal("Not enough arguments.")
	}

	//from the implementation specifications...assumed inclusive
	if nBalls < 27 || nBalls > 127 {
		log.Fatal("Number of balls must be between 27 and 127 (inclusive)")
	}

	nMinutes := 0
	if strings.Compare(options[0], "-mode2") == 0 {
		if len(options) > 2 {
			nMinutes, err = strconv.Atoi(options[2])
			if err != nil {
				log.Fatal("Error converting nMinutes to integer.")
			}
		} else {
			//set default to 1 day.
			fmt.Println("No minutes argument supplied, setting to 1440.")
			nMinutes = 1440
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
		log.Fatal("Unknown option. Must be -mode1 or -mode2")
	}
}
