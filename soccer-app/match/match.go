package match

import (
	"fmt"
	"math/rand"
	m "soccer/model"
	"time"
)

func shuffleTeams(teams []m.Team) [][]m.Team {
	rand.Shuffle(len(teams), func(i, j int) {
		teams[i], teams[j] = teams[j], teams[i]
	})
	var pairs [][]m.Team
	for i := 0; i < len(teams); i += 2 {
		pair := []m.Team{teams[i], teams[i+1]}
		pairs = append(pairs, pair)
	}
	return pairs
}

func makeScores(players []m.Player, maxScore int) (int, []m.Player) {
	randomScoreValue := rand.Intn(maxScore)
	var scoredPlayers []m.Player

	for i := 0; i < randomScoreValue; i++ {
		randomPlayerIndex := rand.Intn(randomScoreValue)
		scoredPlayers = append(scoredPlayers, players[randomPlayerIndex])
	}
	return randomScoreValue, scoredPlayers
}

func penaltySeries(team1, team2 m.Team) (int, m.Team) {
	scores1, scores2 := 0, 0
	players := []m.Player{}

	for i := 0; i < 5; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		randomPlayerFromTeam1 := team1.Players[r.Intn(12)]
		randomPlayerFromTeam2 := team2.Players[r.Intn(12)]

		attempt1 := rand.Intn(2)
		attempt2 := rand.Intn(2)
		scores1 += attempt1
		scores2 += attempt2
		players = append(players, randomPlayerFromTeam1, randomPlayerFromTeam2)
	}

	if scores1 > scores2 {
		fmt.Printf("Penalty series ⚽️ %s - %s (%d:%d) ⚽️\n", team1.Name, team2.Name, scores1, scores2)
		return scores1, team1
	} else if scores1 == scores2 {
		return penaltySeries(team1, team2)
	} else {
		fmt.Printf("Penalty series ⚽️ %s - %s (%d:%d) ⚽️\n", team1.Name, team2.Name, scores1, scores2)
		return scores2, team2
	}
}

func PlayRound(teams []m.Team) []m.Team {
	shuffledTeams := shuffleTeams(teams)
	var winners []m.Team

	for i := 0; i < len(shuffledTeams); i++ {
		winner := simulateMatch(shuffledTeams[i][0], shuffledTeams[i][1])
		winners = append(winners, winner)
	}

	return winners
}

func simulateMatch(team1, team2 m.Team) m.Team {
	players1 := team1.Players
	players2 := team2.Players

	score1, _ := makeScores(players1, 5)
	score2, _ := makeScores(players2, 5)

	fmt.Printf("%s - %s (%d:%d)\n", team1.Name, team2.Name, score1, score2)

	if score1 > score2 {
		return team1
	} else if score1 == score2 {
		_, t := penaltySeries(team1, team2)
		return t
	} else {
		return team2
	}
}
