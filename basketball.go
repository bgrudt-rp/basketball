package main

import (
	"basketball/game"
	"basketball/gameclock"
)

func main() {
	gameclock.ClockTest()
	game.StartGame(1)
}
