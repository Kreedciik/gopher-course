package main

import (
	"fmt"
	f "soccer/format"
	fs "soccer/fs"
	match "soccer/match"
	"time"
)

func main() {
	soccerData := fs.ReadDataFromJSON("soccer.json")
	f.WelcomeText("UEFA Champions League")

	// Winners in 1/16
	fmt.Println("Statistics for 1/16")
	winners1 := match.PlayRound(soccerData)
	fmt.Println("")
	time.Sleep(2 * time.Second)

	// Winners in 1/8
	fmt.Println("Statistics for 1/8")
	winners2 := match.PlayRound(winners1)
	fmt.Println("")
	time.Sleep(2 * time.Second)

	// Winners in 1/4
	fmt.Println("Statistics for 1/4")
	winners3 := match.PlayRound(winners2)
	fmt.Println("")
	time.Sleep(2 * time.Second)

	// Winners in 1/2
	fmt.Println("Statistics for 1/2")
	champion := match.PlayRound(winners3)
	fmt.Println("")
	time.Sleep(2 * time.Second)

	// Champion
	f.PrintWinner(champion[0].Name)
}
