package main

//ClockState -- the current state of the ball clock
type ClockState struct {
	Min     []int //holds 4 balls
	FiveMin []int //holds 11 balls
	Hour    []int //holds 11 balls, but always contains 1 (so count starts @ 1)
	Main    []int //holds at most 'nBalls'
}

//Init -- initializes the state of the clock
func (cs *ClockState) Init(nBalls int) {
	cs.Min = []int{}
	cs.FiveMin = []int{}
	cs.Hour = []int{}
	cs.Main = []int{}

	//setup the main queue, assuming in order starting positions...
	for i := 1; i <= nBalls; i++ {
		cs.Main = append(cs.Main, i)
	}
}

//PopBall -- pops the next ball off of the front the main queue
func (cs *ClockState) PopBall() int {
	nextBall := cs.Main[0]
	cs.Main = RemoveIndex(cs.Main, 0)
	return nextBall
}

//AddMin -- adds a new ball to the minute queue
func (cs *ClockState) AddMin(newBall int) {
	if len(cs.Min) == 4 {
		//clear array, drop current 4 balls back into queue in reverse order
		oldMinQueue := ReverseQueue(cs.Min)
		//clear minute queue
		cs.Min = []int{}
		//append the balls from Min onto Main in reverse order
		cs.Main = append(cs.Main, oldMinQueue...)
		//add new ball to FiveMin
		cs.AddFiveMin(newBall)
	} else {
		cs.Min = append(cs.Min, newBall)
	}
}

//AddFiveMin -- adds a ball to the 5 min queue
func (cs *ClockState) AddFiveMin(newBall int) {
	if len(cs.FiveMin) == 11 {
		oldFiveMinQueue := ReverseQueue(cs.FiveMin)
		cs.FiveMin = []int{}
		cs.Main = append(cs.Main, oldFiveMinQueue...)
		cs.AddHour(newBall)
	} else {
		cs.FiveMin = append(cs.FiveMin, newBall)
	}
}

//AddHour -- adds a ball to the hour queue
func (cs *ClockState) AddHour(newBall int) {
	if len(cs.Hour) == 11 {
		//everything already in the queue goes back to main queue in reverse order
		oldHourQueue := ReverseQueue(cs.Hour)
		cs.Main = append(cs.Main, oldHourQueue...)
		//... before the ball itself returns to the main queue
		cs.Main = append(cs.Main, newBall)
		cs.Hour = []int{}
	} else {
		cs.Hour = append(cs.Hour, newBall)
	}
}

//IncrementByMinute -- increments the state of the clock by 1 minute
func (cs *ClockState) IncrementByMinute() {
	currentBall := cs.PopBall()
	cs.AddMin(currentBall)
}
