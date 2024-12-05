package statistics

import (
	"math/rand"
	m "soccer/model"
)

func SetStatistics(playerStatistics m.PlayerStatistics, scoredPlayers []m.Player) {
	for _, player := range scoredPlayers {
		pScore, ok := playerStatistics[player.Id]
		if ok {
			pScore.Goals += 1
		} else {
			pScore.Goals = 1
		}
		pScore.PlayerName = player.Name
		pScore.Position = player.Position
		playerStatistics[player.Id] = pScore
	}
}

func FindTopScorer(playerStatistics m.PlayerStatistics) m.PlayerScore {
	var topScoredPlayer m.PlayerScore = m.PlayerScore{Goals: 0}

	for _, player := range playerStatistics {
		if topScoredPlayer.Goals < player.Goals {
			topScoredPlayer = player
		} else if topScoredPlayer.Goals == player.Goals {
			randomTopPlayer := rand.Intn(2)
			if randomTopPlayer == 0 {
				topScoredPlayer = player
			}
		}
	}

	return topScoredPlayer
}
