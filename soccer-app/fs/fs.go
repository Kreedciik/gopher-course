package fs

import (
	"encoding/json"
	"os"
	m "soccer/model"
)

func ReadDataFromJSON(fileName string) []m.Team {
	data, err := os.ReadFile(fileName)
	var soccerTeams []m.Team
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &soccerTeams)
	if err != nil {
		panic(err)
	}

	return soccerTeams
}
