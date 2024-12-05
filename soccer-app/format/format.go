package format

import (
	"fmt"
	m "soccer/model"
	"strings"
)

func PrintTeams(data []m.Team) {
	for _, team := range data {
		fmt.Println(team.Name)
	}
}

func print(text string) {
	width := 50 // Total width of the display

	// Trophy text with padding
	// trophyText := fmt.Sprintf("🏆 %s 🏆", teamName)
	padding := (width - len(text)) / 2
	centeredText := fmt.Sprintf("%s%s%s", strings.Repeat(" ", padding), text, strings.Repeat(" ", padding))

	// Top and bottom decorative lines
	line := strings.Repeat("*", width)

	// Print the trophy display
	fmt.Println(line)
	fmt.Println(centeredText)
	fmt.Println(line)
}

func PrintWinner(name string) {
	print(fmt.Sprintf("🏆 %s 🏆", name))
}

func WelcomeText(text string) {
	print(fmt.Sprintf("⚽️ %s ⚽️", text))
	fmt.Println()
}
