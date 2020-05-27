package game

import (
	"basketball/data"
	"basketball/gameclock"
	"fmt"
)

//ShotList - A list of shots in a game
var ShotList []Shot
var c gameclock.GameClock

//StartGame - Designates the beginning of an active game
func StartGame(a int) {
	t := make(map[string]data.Team)
	t["home"], t["away"] = getTeamsByGameID(a)
	c = gameclock.NewClock(a)
	liveGame(t)
	endGame(a)
}

//endGame - Process steps to complete an active game
func endGame(i int) {
	printGameScore(i)
}

//printGameScore - Given a game ID, print a score.
func printGameScore(i int) {
	var h int
	var a int
	t := make(map[string]data.Team)
	m := make(map[int]int)
	for _, s := range ShotList {
		if s.madeShot {
			for _, c := range data.TeamList {
				for _, p := range c.Players {
					if s.shooter == p.PlayerID {
						m[c.TeamID] += s.value
					}
				}
			}
		}
	}

	t["home"], t["away"] = getTeamsByGameID(i)
	for k, v := range m {
		if k == t["home"].TeamID {
			h = v
		}
		if k == t["away"].TeamID {
			a = v
		}
	}
	fmt.Printf("\n\nFinal Score:\n"+
		"%s : %d\n"+
		"%s : %d\n", t["home"].Name, h, t["away"].Name, a)
}
