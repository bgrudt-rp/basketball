package gameclock

import (
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

//GameTime input of time.Duration returns a quarter and time remaining value
func GameTime(d time.Duration) (int, int, int) {
	i := int(d)
	q := i / 720000000000
	q++
	t := (720000000000 * q) - i
	m := t / 60000000000
	s := (t % 60000000000) / 1000000000
	return q, m, s
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

//TimeElapsed - show time elapsed in seconds
func TimeElapsed(c GameClock) time.Duration {
	var td time.Duration
	if c.currStart.Valid {
		t := time.Now()
		td = td + t.Sub(c.currStart.Time)
	}
	for _, s := range c.gameDuration {
		td = td + s.incDuration
	}
	return td
}
