package data

//Game - struct to be  used to find a specific game
type Game struct {
	GameID   int
	HomeTeam int
	AwayTeam int
}

//Player - struct to be stored in Team
type Player struct {
	PlayerID int
	Name     string
	Number   int
}

//Team - struct to define all teams
type Team struct {
	TeamID  int
	City    string
	Name    string
	Players []Player
}

//This data is a mock up of a database

//GameList - list of all games played in our season
var GameList = []Game{
	{
		GameID:   1,
		HomeTeam: 1,
		AwayTeam: 2,
	},
	{
		GameID:   2,
		HomeTeam: 2,
		AwayTeam: 3,
	},
}

//TeamList - let's build a few teams into our league
var TeamList = []Team{
	{
		TeamID: 1,
		City:   "Detroit",
		Name:   "Pistons",
		Players: []Player{
			{
				PlayerID: 1,
				Name:     "Isiah Thomas",
				Number:   11,
			},
			{
				PlayerID: 2,
				Name:     "Bill Laimbeer",
				Number:   40,
			},
			{
				PlayerID: 3,
				Name:     "Vinnie Johnson",
				Number:   15,
			},
			{
				PlayerID: 4,
				Name:     "John Salley",
				Number:   22,
			},
			{
				PlayerID: 5,
				Name:     "Joe Dumars",
				Number:   4,
			},
		},
	},
	{
		TeamID: 2,
		City:   "Chicago",
		Name:   "Bulls",
		Players: []Player{
			{
				PlayerID: 6,
				Name:     "Michael Jordan",
				Number:   23,
			},
			{
				PlayerID: 7,
				Name:     "Scottie Pippen",
				Number:   33,
			},
			{
				PlayerID: 8,
				Name:     "Dennis Rodman",
				Number:   12,
			},
			{
				PlayerID: 9,
				Name:     "Luc Longley",
				Number:   13,
			},
			{
				PlayerID: 10,
				Name:     "Ron Harper",
				Number:   9,
			},
		},
	},
	{
		TeamID: 3,
		City:   "Los Angeles",
		Name:   "Lakers",
		Players: []Player{
			{
				PlayerID: 11,
				Name:     "James Worthy",
				Number:   42,
			},
			{
				PlayerID: 12,
				Name:     "Magic Johnson",
				Number:   32,
			},
			{
				PlayerID: 13,
				Name:     "Sam Perkins",
				Number:   14,
			},
			{
				PlayerID: 14,
				Name:     "Byron Scott",
				Number:   4,
			},
			{
				PlayerID: 15,
				Name:     "Vlade Divac",
				Number:   12,
			},
		},
	},
}
