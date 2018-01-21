package main

import "testing"

func TestParseArgs(t *testing.T) {
	var options1 = []string{"-mode1", "30"}
	nBalls, nMinutes := parseArgs(options1)
	if nBalls != 30 {
		t.Errorf("nBalls != %d instead is %d", 30, nBalls)
	}
	if nMinutes != 0 {
		t.Errorf("nMinutes != %d instead is %d", 0, nMinutes)
	}

	var options2 = []string{"-mode2", "100", "300"}
	nBalls, nMinutes = parseArgs(options2)
	if nBalls != 100 {
		t.Errorf("nBalls != %d instead is %d", 100, nBalls)
	}
	if nMinutes != 300 {
		t.Errorf("nMinutes != %d instead is %d", 300, nMinutes)
	}
}

func TestInitState(t *testing.T) {
	nBalls := 30
	initState := initClockState(nBalls)
	//check sizes
	if len(initState.Main) != nBalls {
		t.Errorf("Main queue was incorrect, got: %d, want: %d.", len(initState.Main), nBalls)
	}
	//should be in order initially
	for i := 1; i <= nBalls; i++ {
		if initState.Main[i-1] != i {
			t.Errorf("initState[%d] is %d and not %d", i, initState.Main[i], i)
		}
	}
}

func TestRemoveBall(t *testing.T) {
	nBalls := 27
	state := initClockState(nBalls)
	currentBall := state.Main[0]
	state.Main = RemoveBall(state.Main, 0)
	if currentBall != 1 {
		t.Errorf("currentBall != %d but is instead %d", 0, currentBall)
	}
	if len(state.Main) != 26 {
		t.Errorf("length of main queue is %d", len(state.Main))
	}
}

func TestAddMin(t *testing.T) {
	nBalls := 27
	state := initClockState(nBalls)
	currentBall := state.PopBall()
	state.AddMin(currentBall)

	if len(state.Min) != 1 {
		t.Errorf("minutes queue length is %d and not 1", len(state.Min))
	}
}
func TestAddFiveMin(t *testing.T) {
	nBalls := 27
	state := initClockState(nBalls)
	for i := 0; i < 5; i++ {
		currentBall := state.PopBall()
		state.AddMin(currentBall)
	}
	if len(state.Min) != 0 {
		t.Errorf("Length of minutes queue is %d and not 0", len(state.Min))
	}
	if len(state.FiveMin) != 1 {
		t.Errorf("Length of five minutes queue is %d and not 1", len(state.FiveMin))
	}
}
func TestAddHour(t *testing.T) {
	nBalls := 27
	state := initClockState(nBalls)
	for i := 0; i < 60; i++ {
		currentBall := state.PopBall()
		state.AddMin(currentBall)
	}
	if len(state.Min) != 0 {
		t.Errorf("Length of minutes queue is %d and not 0", len(state.Min))
	}
	if len(state.FiveMin) != 0 {
		t.Errorf("Length of five minutes queue is %d and not 0", len(state.FiveMin))
	}
	if len(state.Hour) != 1 {
		t.Errorf("Length of hour queue is %d and not 1", len(state.Hour))
	}
}

func TestFullDay(t *testing.T) {
	nBalls := 27
	state := initClockState(nBalls)
	for i := 0; i < 1440; i++ {
		currentBall := state.PopBall()
		state.AddMin(currentBall)
	}
	if len(state.Min) != 0 {
		t.Errorf("Length of minutes queue is %d and not 0", len(state.Min))
	}
	if len(state.FiveMin) != 0 {
		t.Errorf("Length of five minutes queue is %d and not 0", len(state.FiveMin))
	}
	if len(state.Hour) != 0 {
		t.Errorf("Length of hour queue is %d and not 0", len(state.Hour))
	}
}

func TestMode1(t *testing.T) {
	//var options1 = []string{"-mode1", "30"}
	//var options2 = []string{"-mode1", "45"}
}

func TestMode2(t *testing.T) {
	//var options = []string{"-mode2", "30", "325"}
}
