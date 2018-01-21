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

func TestAddMin(t *testing.T) {
	nBalls := 30
	state := initClockState(nBalls)
	currentBall, newMain := RemoveIndex(state.Main, 0)
	state.Main = newMain
	state.AddMin(currentBall)
}
func TestIncrementFiveMinutes(t *testing.T) {
}
func TestIncrementHours(t *testing.T) {
}

func TestMode1(t *testing.T) {
	//var options1 = []string{"-mode1", "30"}
	//var options2 = []string{"-mode1", "45"}
}

func TestMode2(t *testing.T) {
	//var options = []string{"-mode2", "30", "325"}
}
