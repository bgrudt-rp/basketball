package game

import (
	"basketball/data"
)

//ShotList - A list of shots in a game
var ShotList []Shot

//StartGame - Designates the beginning of an active game
func StartGame(a int) {
	t := make(map[string]data.Team)
	t["home"], t["away"] = getTeamsByGameID(a)
	liveGame(t)
	endGame(a)
}

func endGame(i int) {
	printGameScore(i)
}
