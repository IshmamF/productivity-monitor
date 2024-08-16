package display

import (
	"github.com/IshmamF/productivity-monitor/database"
	"github.com/IshmamF/productivity-monitor/utils"
	"fmt"
)

func (t *Display) WeeklyDisplay(db *database.DB) (string) {
	currentTime := utils.GetCurrentTimestamp()
	secondsInDay := int64(604800)
	data := db.CountAppUsageWithRange(currentTime - secondsInDay, currentTime)

	bar := ConvertToBarList(data)
	
	ShowBar(bar, "weekly")
	var input string
	ShowOptions()
	fmt.Scan(&input)

	return HandleWeeklyInput(input)
}

func HandleWeeklyInput(input string) (selectedOption string) {
	switch input {
	case "p":
		selectedOption = "Statistics"
	case "m": 
		selectedOption = "Menu"
	default:
		selectedOption = "Weekly"
	}
	return
}