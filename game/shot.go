package game

import (
	"basketball/data"
	"fmt"
)

//Shot - EVENT - Player attempted shot
type Shot struct {
	value     int
	shooter   int
	passer    int
	rebounder int
	madeShot  bool
}

//ShotAttempt - takes a ShotList slice and adds a new Shot struct
func ShotAttempt(value, shooter, passer, rebounder int, madeShot bool, ShotList *[]Shot, t map[string]data.Team) {
	var s Shot
	s.value = value
	s.shooter = shooter
	s.passer = passer
	s.rebounder = rebounder
	s.madeShot = madeShot
	*ShotList = append(*ShotList, s)
	if s.madeShot {
		shotOutputMade(s, t)
	} else {
		shotOutputMiss(s, t)
	}
}

func shotOutputMade(s Shot, t map[string]data.Team) {
	var sh data.Player
	var a data.Player
	sType := shotType(s.value)
	sh = getPlayerByID(s.shooter, t)
	a = getPlayerByID(s.passer, t)
	fmt.Printf("\nSHOT WAS MADE\n"+
		"%s hit a %s. ", sh.Name, sType)
	if s.passer > 0 {
		fmt.Printf("Assisted by %s.\n", a.Name)
	}
}
func shotOutputMiss(s Shot, t map[string]data.Team) {
	var sh data.Player
	sType := shotType(s.value)
	sh = getPlayerByID(s.shooter, t)

	fmt.Printf("\nMissed shot\n"+
		"%s missed a %s.\n", sh.Name, sType)
}

func shotType(a int) string {
	var sType string
	switch a {
	case 1:
		sType = "free throw"
	case 2:
		sType = "two pointer"
	case 3:
		sType = "three"
	}
	return sType
}
