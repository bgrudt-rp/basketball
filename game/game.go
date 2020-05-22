package game

import (
	"fmt"
)

//Team - struct to be stored in TeamList
type Team struct {
	teamID   int
	city     string
	name     string
	color    string
	teamType string
}

//TeamScore - struct to be stored in GameScore
type TeamScore struct {
	teamID    int
	teamScore int
}

//Player - struct to be stored in PlayerList
type Player struct {
	playerID int
	name     string
	number   int
	team     int
}

//ShotList - A list of shots in a game
var ShotList []Shot

//TeamList - list of teams in game
var TeamList []Team

//PlayerList - list of players in game
var PlayerList []Player

//GameScore - score of game
var GameScore []TeamScore

//StartGame - Designates the beginning of an active game
func StartGame() {
	zeroScore()
	addTeam(1, "Detroit", "Pistons", "Blue", "Away", &TeamList)
	addTeam(2, "Chicago", "Bulls", "White", "Home", &TeamList)

	addPlayer(1, "Isiah Thomas", 11, 1, &PlayerList)
	addPlayer(2, "Bill Laimbeer", 40, 1, &PlayerList)
	addPlayer(3, "Vinnie Johnson", 15, 1, &PlayerList)
	addPlayer(4, "John Salley", 22, 1, &PlayerList)
	addPlayer(5, "Joe Dumars", 4, 1, &PlayerList)
	addPlayer(6, "Michael Jordan", 23, 2, &PlayerList)
	addPlayer(7, "Scottie Pippen", 33, 2, &PlayerList)
	addPlayer(8, "Dennis Rodman", 12, 2, &PlayerList)
	addPlayer(9, "Luc Longley", 13, 2, &PlayerList)
	addPlayer(10, "Ron Harper", 9, 2, &PlayerList)

	gameStuff()
	endGame()
}
func addPlayer(a int, b string, c, d int, PlayerList *[]Player) {
	var p Player
	p.playerID = a
	p.name = b
	p.number = c
	p.team = d
	*PlayerList = append(*PlayerList, p)
}
func addTeam(a int, b, c, d, e string, TeamList *[]Team) {
	var t Team
	t.teamID = a
	t.city = b
	t.name = c
	t.color = d
	t.teamType = e
	*TeamList = append(*TeamList, t)
}

func gameStuff() {

	ShotAttempt(2, 1, 2, -1, true, &ShotList)
	ShotList = append(ShotList, ShotAttempt2(3, 4, 3, 9, false))
	ShotAttempt(2, 6, 7, -1, true, &ShotList)
	ShotAttempt(1, 6, 7, 2, false, &ShotList)
	ShotAttempt(3, 5, -1, -1, true, &ShotList)
}

func endGame() {
	Team1Score, Team1Name, Team2Score, Team2Name := getTeamScores()
	fmt.Printf("\n\nFinal Score:\n"+
		"%s: %d\n"+
		"%s : %d\n", Team1Name, Team1Score, Team2Name, Team2Score)
}
func zeroScore() {
	var s TeamScore
	s.teamID = 1
	s.teamScore = 0
	GameScore = append(GameScore, s)
	s.teamID = 2
	s.teamScore = 0
	GameScore = append(GameScore, s)
}
