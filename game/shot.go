package game

import (
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

func scoreUpdate(a, b int, GameScore []TeamScore) {
	p := getPlayerByID(b)
	t := getTeamByPlayerID(p.team)
	tid := t.teamID
	var cScore int
	var loc int
	for i, t := range GameScore {
		if t.teamID == tid {
			cScore = t.teamScore
			loc = i
		}
	}
	nScore := cScore + a
	GameScore[loc].teamScore = nScore
}

//ShotAttempt - takes a ShotList slice and adds a new Shot struct
func ShotAttempt(value, shooter, passer, rebounder int, madeShot bool, ShotList *[]Shot) {
	var s Shot
	s.value = value
	s.shooter = shooter
	s.passer = passer
	s.rebounder = rebounder
	s.madeShot = madeShot
	*ShotList = append(*ShotList, s)
	if s.madeShot {
		scoreUpdate(s.value, s.shooter, GameScore)
		shotOutputMade(s)
	} else {
		shotOutputMiss(s)
	}
}

//ShotAttempt2 - returns a Shot struct to caller
func ShotAttempt2(value, shooter, passer, rebounder int, madeShot bool) Shot {
	var s Shot
	s.value = value
	s.shooter = shooter
	s.passer = passer
	s.rebounder = rebounder
	s.madeShot = madeShot
	if s.madeShot {
		scoreUpdate(s.value, s.shooter, GameScore)
		shotOutputMade(s)
	} else {
		shotOutputMiss(s)
	}
	return s
}
func shotOutputMade(s Shot) {
	var sh Player
	var a Player
	sType := shotType(s.value)
	sh = getPlayerByID(s.shooter)
	a = getPlayerByID(s.passer)
	fmt.Printf("\nSHOT WAS MADE\n"+
		"%s hit a %s. ", sh.name, sType)
	if s.passer > 0 {
		fmt.Printf("Assisted by %s.\n", a.name)
	}
}
func shotOutputMiss(s Shot) {
	var sh Player
	sType := shotType(s.value)
	sh = getPlayerByID(s.shooter)

	fmt.Printf("\nMissed shot\n"+
		"%s missed a %s.\n", sh.name, sType)
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
