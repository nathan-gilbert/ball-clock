package main

import "testing"

func TestInitState(t *testing.T) {
	nBalls := 30
	initState := initClockState(nBalls)
	//check sizes
	if len(initState.Main) != nBalls {
		t.Errorf("Main queue was incorrect, got: %d, want: %d.", len(initState.Main), nBalls)
	}
	//should be in order initially
	for i := 1; i <= nBalls; i++ {
		if initState.Main[i] != i {
			t.Errorf("initState[%d] != %d", i, i)
		}
	}
}
