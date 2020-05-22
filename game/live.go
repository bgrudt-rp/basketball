package game

import "basketball/data"

func liveGame(t map[string]data.Team) {

	ShotAttempt(2, 1, 2, -1, true, &ShotList, t)
	ShotAttempt(2, 6, 7, -1, true, &ShotList, t)
	ShotAttempt(1, 6, 7, 2, false, &ShotList, t)
	ShotAttempt(3, 5, -1, -1, true, &ShotList, t)
}
