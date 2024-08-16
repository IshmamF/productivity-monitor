package display

import (
	"github.com/IshmamF/productivity-monitor/database"
	"github.com/IshmamF/productivity-monitor/utils"
	"math"
	"github.com/pterm/pterm"
	"fmt"
)

func (t *Display) DailyDisplay(db *database.DB) (string) {
	currentTime := utils.GetCurrentTimestamp()
	secondsInDay := int64(86400)
	data := db.CountAppUsageWithRange(currentTime - secondsInDay, currentTime)

	bar := ConvertToBarList(data)
	
	ShowDailyBar(bar)
	var input string
	ShowOptions()
	fmt.Scan(&input)

	return HandleDailyInput(input)
}

func ConvertToBarList(data []database.App_Count) (bar []pterm.Bar) {
	for i, activity := range data {
		minutes := int(math.Ceil((float64(activity.Count) / float64(60))))
		style := ChooseStyle(i)
		bar = append(bar, pterm.Bar{Label: activity.App_Name, Value: minutes, Style: style})
	}
	return
}

func HandleDailyInput(input string) (selectedOption string) {
	switch input {
	case "p":
		selectedOption = "Statistics"
	case "m": 
		selectedOption = "Menu"
	default:
		selectedOption = "Daily"
	}
	return
}

func ShowDailyBar(bar []pterm.Bar) {
	pterm.DefaultBasicText.Println("")
	pterm.DefaultBasicText.Println("Most Used Apps within 24 Hours in Minutes")
	pterm.DefaultBarChart.WithBars(bar).WithHorizontal().WithShowValue().WithWidth(40).Render()
}

func ChooseStyle(i int) (style *pterm.Style) {
	if i % 10 == 0 {
		style = pterm.NewStyle(pterm.FgGreen)
	} else if i % 10 == 1 {
		style = pterm.NewStyle(pterm.FgBlue)
	} else  if i % 10 == 2{
		style = pterm.NewStyle(pterm.FgLightRed)
	} else if i % 10 == 3 {
		style = pterm.NewStyle(pterm.FgLightCyan)
	} else if i % 10 == 4 {
		style = pterm.NewStyle(pterm.FgLightMagenta)
	} else if i % 10 == 5 {
		style = pterm.NewStyle(pterm.FgLightYellow)
	} else if i % 10 == 6 {
		style = pterm.NewStyle(pterm.FgRed)
	} else if i % 10 == 7 {
		style = pterm.NewStyle(pterm.FgLightGreen)
	} else if i % 10 == 8 {
		style = pterm.NewStyle(pterm.FgMagenta)
	} else {
		style = pterm.NewStyle(pterm.FgGray)
	}
	return
}