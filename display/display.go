package display

import (
	"strconv"
	"strings"
	"github.com/IshmamF/productivity-monitor/database"
	"github.com/IshmamF/productivity-monitor/utils"
	"github.com/pterm/pterm"
	"fmt"
	"math"
)

type Display struct {
	tracking_options_on []string
	tracking_options_off []string
	menu_options []string
	alert_on_options []string
	alert_off_options [] string
	statistic_options []string 
}

func (t *Display) Init() {
	t.menu_options = append(t.menu_options, "Track Activity")
	t.menu_options = append(t.menu_options, "Statistics")
	t.menu_options = append(t.menu_options, "Alert Settings")
	t.menu_options = append(t.menu_options, "Quit Program")

	t.tracking_options_off = append(t.tracking_options_off, "Start Tracking")
	t.tracking_options_off = append(t.tracking_options_off, "Menu")

	t.tracking_options_on = append(t.tracking_options_on, "Stop Tracking")
	t.tracking_options_on = append(t.tracking_options_on, "View Current Session Data")
	t.tracking_options_on = append(t.tracking_options_on, "Menu")

	t.alert_off_options = append(t.alert_off_options, "Turn Alerts On")
	t.alert_off_options = append(t.alert_off_options, "Menu")

	t.alert_on_options = append(t.alert_on_options, "Turn Alerts Off")
	t.alert_on_options = append(t.alert_on_options, "Set Interval")
	t.alert_on_options = append(t.alert_on_options, "Menu")

	t.statistic_options = append(t.statistic_options, "Daily")
	t.statistic_options = append(t.statistic_options, "Weekly")
	t.statistic_options = append(t.statistic_options, "All Time")
}

func (t *Display) MenuDisplay() (selectedOption string) {
	selectedOption, _ = pterm.DefaultInteractiveSelect.WithOptions(t.menu_options).Show()
	return
}

func (t *Display) TrackingDisplay(choice chan string, running *bool) (selectedOption string) {
	if *running {
		selectedOption, _ = pterm.DefaultInteractiveSelect.WithOptions(t.tracking_options_on).Show()
	} else {
		selectedOption, _ = pterm.DefaultInteractiveSelect.WithOptions(t.tracking_options_off).Show()
	}

	switch selectedOption {
	case "Start Tracking":
		choice <- "start"
	case "Stop Tracking":
		choice <- "stop"
	}
	return
}

func (t *Display) AlertSettingsDisplay(db *database.DB) (selectedOption string) {
	alert_settings := db.GetAlertSettings()
	var options []string
	if alert_settings.Alert_On {
		options = t.alert_on_options
	} else {
		options = t.alert_off_options
	}
	selectedOption, _ = pterm.DefaultInteractiveSelect.WithOptions(options).Show()
	if selectedOption == "Turn Alerts Off" {
		db.UpdateAlertOn(false)
	} else if selectedOption == "Turn Alerts On" {
		db.UpdateAlertOn(true)
	}
	return 
}

func (t *Display) IntervalDisplay(db *database.DB, selectedOption string) (string) {
	alert_settings := db.GetAlertSettings()
	if selectedOption == "Interval Error(1)" {
		pterm.Warning.Println("Invalid Interval: Not an Integer")
	} else if selectedOption == "Interval Error(2)" {
		pterm.Warning.Println("Invalid Interval: Interval Must be Minimum 60 Seconds")
	}
	if alert_settings.Interval == 300 {
		pterm.Info.Println("Current Monitor Interval: ", alert_settings.Interval, " (Default)")
	} else {
		pterm.Info.Println("Current Monitor Interval: ", alert_settings.Interval)
	}
	pterm.Info.Println("Enter -1 to Return to Main Menu")
	TextPrinter := pterm.InteractiveTextInputPrinter{
		TextStyle: pterm.NewStyle(pterm.FgLightGreen),
		DefaultText: "Input Interval in Seconds: ",
	}
	resultStr, _ := TextPrinter.Show()
	resultInt, err := strconv.Atoi(resultStr)
	if err != nil {
		checkErrorString := err.Error()
		if strings.Contains(checkErrorString, "invalid syntax") {
			return "Interval Error(1)"
		} 
    }
	if resultInt != -1 && resultInt < 60 {
		return "Interval Error(2)"
	}
	if resultInt == -1 {
		return "Menu"
	}
	db.UpdateAlertInterval(resultInt)
	return "Menu"
}

func (t *Display) SessionDisplay(db *database.DB, startTime *int64) (selectedOption string) {
	currentTime := utils.GetCurrentTimestamp()
	var tablerow [][]string
	data := db.CountAppUsageWithRange(*startTime, currentTime)
	tablerow = append(tablerow, []string{"App or Site Name", "Usage"})
	for _, activity := range data {
		tablerow = append(tablerow, []string{activity.App_Name, fmt.Sprint(activity.Count)})
	}
	pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tablerow).Render()
	var input string
	pterm.Info.WithMessageStyle(pterm.NewStyle(pterm.FgLightMagenta)).Println("Enter 'p' to Return to Previous Page")
	pterm.Info.WithMessageStyle(pterm.NewStyle(pterm.FgLightMagenta)).Println("Enter 'm' to Return to go to Menu Page")
	pterm.Info.WithMessageStyle(pterm.NewStyle(pterm.FgLightMagenta)).Println("Enter any character to Refresh")
	textprinter := pterm.BasicTextPrinter{
		Style: pterm.NewStyle(pterm.FgLightGreen),
	}
	textprinter.Print("Input character: ")
	fmt.Scan(&input)
	switch input {
	case "p":
		selectedOption = "Track Activity"
	case "m": 
		selectedOption = "Menu"
	default:
		selectedOption = "View Current Session Data"
	}
	return
}

func (t *Display) StatisticsDisplay(db *database.DB) (selectedOption string) {
	selectedOption, _ = pterm.DefaultInteractiveSelect.WithOptions(t.statistic_options).Show()
	return 
}

func (t *Display) DailyDisplay(db *database.DB) (selectedOption string) {
	currentTime := utils.GetCurrentTimestamp()
	secondsInDay := int64(86400)
	data := db.CountAppUsageWithRange(currentTime - secondsInDay, currentTime)
	bar := []pterm.Bar{}
	for i, activity := range data {
		minutes := int(math.Ceil((float64(activity.Count) / float64(60))))
		var style *pterm.Style
		if i % 10 == 0 {
			style = pterm.NewStyle(pterm.FgGreen)
		} else if i % 2 == 0 {
			style = pterm.NewStyle(pterm.FgBlue)
		} else  if i % 3 == 0{
			style = pterm.NewStyle(pterm.FgLightRed)
		} else if i % 4 == 0 {
			style = pterm.NewStyle(pterm.FgCyan)
		} else if i % 5 == 0 {
			style = pterm.NewStyle(pterm.FgCyan)
		} else if i % 6 == 0 {
			style = pterm.NewStyle(pterm.FgCyan)
		} else if i % 7 == 0 {
			style = pterm.NewStyle(pterm.FgCyan)
		} else if i % 8 == 0 {
			style = pterm.NewStyle(pterm.FgCyan)
		}
		bar = append(bar, pterm.Bar{Label: activity.App_Name, Value: minutes, Style: style})
	}
	pterm.DefaultBasicText.Println("")
	pterm.DefaultBasicText.Println("Most Used Apps within 24 Hours in Minutes")
	pterm.DefaultBarChart.WithBars(bar).WithHorizontal().WithShowValue().WithWidth(40).Render()
	var input string
	pterm.Info.WithMessageStyle(pterm.NewStyle(pterm.FgLightMagenta)).Println("Enter 'p' to Return to Previous Page")
	pterm.Info.WithMessageStyle(pterm.NewStyle(pterm.FgLightMagenta)).Println("Enter 'm' to Return to go to Menu Page")
	pterm.Info.WithMessageStyle(pterm.NewStyle(pterm.FgLightMagenta)).Println("Enter any character to Refresh")
	textprinter := pterm.BasicTextPrinter{
		Style: pterm.NewStyle(pterm.FgLightGreen),
	}
	textprinter.Print("Input character: ")
	fmt.Scan(&input)
	switch input {
	case "p":
		selectedOption = "Track Activity"
	case "m": 
		selectedOption = "Menu"
	default:
		selectedOption = "View Current Session Data"
	}
	return
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