package game

func getPlayerByID(i int) Player {
	var pl Player
	for _, p := range PlayerList {
		if p.playerID == i {
			pl = p
		}
	}
	return pl
}
func getTeamByPlayerID(i int) Team {
	var t Team
	for _, p := range TeamList {
		if p.teamID == i {
			t = p
		}
	}
	return t
}
func getTeamScores() (int, string, int, string) {
	var t1 int
	var t2 int
	var n1 string
	var n2 string
	for _, t := range GameScore {
		if t.teamID == 1 {
			t1 = t.teamScore
		}
		if t.teamID == 2 {
			t2 = t.teamScore
		}
	}
	for _, t := range TeamList {
		if t.teamID == 1 {
			n1 = t.city
		}
		if t.teamID == 2 {
			n2 = t.city
		}
	}
	return t1, n1, t2, n2
}
