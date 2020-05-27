package game

import "basketball/data"

func liveGame(t map[string]data.Team) {

	newShotAttempt(2, 1, 2, -1, true, &ShotList, t)
	newShotAttempt(2, 6, 7, -1, true, &ShotList, t)
	newShotAttempt(1, 6, 7, 2, false, &ShotList, t)
	newShotAttempt(3, 5, -1, -1, true, &ShotList, t)
}
