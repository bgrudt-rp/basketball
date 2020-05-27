package game

import (
	"basketball/data"
	"basketball/gameclock"
	"time"
)

func liveGame(t map[string]data.Team) {
	c = gameclock.StartClock(1, c)
	time.Sleep(2 * time.Second)
	newShotAttempt(2, 1, 2, -1, true, &ShotList, t, c)
	time.Sleep(4 * time.Second)
	newShotAttempt(2, 6, 7, -1, true, &ShotList, t, c)
	time.Sleep(1 * time.Second)
	c = gameclock.StopClock(1, c)
	newShotAttempt(1, 6, 7, 2, false, &ShotList, t, c)
	c = gameclock.StartClock(1, c)
	time.Sleep(1 * time.Second)
	c = gameclock.StopClock(1, c)
	c = gameclock.AdjustClock(1, c, 2.0)
	time.Sleep(1 * time.Second)
	newShotAttempt(3, 5, -1, -1, true, &ShotList, t, c)
}
