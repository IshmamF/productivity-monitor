package display

import (
	"github.com/IshmamF/productivity-monitor/database"
	"fmt"
)

func (t *Display) AllTimeDisplay(db *database.DB) (string) {
	data := db.AllTimeMostUsedApp()

	bar := ConvertToBarList(data)
	
	ShowBar(bar, "all time")
	var input string
	ShowOptions()
	fmt.Scan(&input)

	return HandleWeeklyInput(input)
}

func HandleAllTimeInput(input string) (selectedOption string) {
	switch input {
	case "p":
		selectedOption = "Statistics"
	case "m": 
		selectedOption = "Menu"
	default:
		selectedOption = "All Time"
	}
	return
}