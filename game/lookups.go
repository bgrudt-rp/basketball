package game

import (
	"basketball/data"
)

//getPlayerByID - returns a data.Player value given an ID input
func getPlayerByID(p int, t map[string]data.Team) data.Player {
	var pl data.Player
	for _, l := range t {
		for _, il := range l.Players {
			if il.PlayerID == p {
				pl = il
			}
		}
	}
	return pl
}

//getTeamsByGame - returns two data.Team values for a given game ID
func getTeamsByGameID(i int) (data.Team, data.Team) {
	var h, a data.Team
	for _, g := range data.GameList {
		if g.GameID == i {
			for _, c := range data.TeamList {
				if c.TeamID == g.HomeTeam {
					h = c
				}
				if c.TeamID == g.AwayTeam {
					a = c
				}
			}
		}
	}

	return h, a
}
