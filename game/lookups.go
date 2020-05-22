package game

import (
	"basketball/data"
	"fmt"
)

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
