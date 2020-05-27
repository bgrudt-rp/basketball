package gameclock

import (
	"fmt"
	"time"
)

//GameClock - parent element for running game cloc
type GameClock struct {
	gameID       int
	currStart    NullTime
	gameDuration []GameInc
}

//GameInc - contains slice of all game increments
type GameInc struct {
	incStart    time.Time
	incStop     time.Time
	incType     string
	incDuration time.Duration
}

//NullTime - allows for exception handling for starting/stopping
type NullTime struct {
	Time  time.Time
	Valid bool
}

//ClockTest - temporary function to test clock
func ClockTest() {
	currGameClock := NewClock(1)
	currGameClock = StartClock(1, currGameClock)
	time.Sleep(2 * time.Second)
	currGameClock = StopClock(1, currGameClock)
	currGameClock = StartClock(1, currGameClock)
	time.Sleep(1 * time.Second)
	currGameClock = StopClock(1, currGameClock)
	currGameClock = AdjustClock(1, currGameClock, 2.0)
	timeElapsed(currGameClock)
}

//NewClock - generates a variable of type GameClock - beginning of game
func NewClock(g int) GameClock {
	var c GameClock
	c.gameID = g

	return c
}

//AdjustClock - manually adjust clock - duration value in seconds
func AdjustClock(g int, c GameClock, t time.Duration) GameClock {
	n := time.Now()
	var cd GameInc
	cd.incStart = n
	cd.incStop = n
	cd.incType = "Time Adjustment"
	cd.incDuration = t * 1000000000
	c.gameDuration = append(c.gameDuration, cd)
	return c
}

//StartClock - start running game clock
func StartClock(g int, c GameClock) GameClock {
	if !c.currStart.Valid {
		t := time.Now()
		c.currStart.Valid = true
		c.currStart.Time = t
	}
	return c
}

//StopClock - stop running game clock for stoppage in play
func StopClock(g int, c GameClock) GameClock {
	if c.currStart.Valid {
		t := time.Now()
		var cd GameInc
		c.currStart.Valid = false
		cd.incStart = c.currStart.Time
		cd.incStop = t
		cd.incType = "Game Increment - Stopped"
		cd.incDuration = t.Sub(c.currStart.Time)
		c.gameDuration = append(c.gameDuration, cd)
	}
	return c
}

func timeElapsed(c GameClock) {
	var td time.Duration
	for _, s := range c.gameDuration {
		td = td + s.incDuration
	}
	fmt.Printf("Total time elapsed in seconds: %v\n", td.Seconds())
}
