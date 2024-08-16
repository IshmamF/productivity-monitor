package main

import (
	_"fmt"
	"github.com/IshmamF/productivity-monitor/darwin"
	_"github.com/IshmamF/productivity-monitor/utils"
	"github.com/IshmamF/productivity-monitor/database"
	"github.com/IshmamF/productivity-monitor/display"
    "github.com/inancgumus/screen"
	"github.com/pterm/pterm"
	"strings"
	"os"
	"time"
)

var (
	db = &database.DB{}
	t = &display.Display{}
	selectedOption string
	startTime int64
	running = false
)

func main() {
	db.Connection()
	ticker := time.NewTicker(time.Second)

	t.Init()

	choice := make(chan string)
	go darwin.Start_Tracking(choice, db, &startTime, &running, ticker)

	for {
		screen.Clear()
		screen.MoveTopLeft()
		pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgBlack)).WithTextStyle(pterm.NewStyle(pterm.FgLightCyan)).Println("Productivity Monitor")

		if strings.Contains(strings.ToLower(selectedOption),"track") {
			selectedOption = t.TrackingDisplay(choice, &running)
		} else if strings.Contains(strings.ToLower(selectedOption),"alert")  {
			selectedOption = t.AlertSettingsDisplay(db)
		} else if strings.Contains(strings.ToLower(selectedOption), "interval") {
			selectedOption = t.IntervalDisplay(db, selectedOption)
		} else if selectedOption == "View Current Session Data" {
			selectedOption = t.SessionDisplay(db, &startTime)
		} else if selectedOption == "Quit Program" {
			os.Exit(1)
		} else if strings.Contains(strings.ToLower(selectedOption),"statistics") {
			selectedOption = t.StatisticsDisplay(db)
		} else if selectedOption == "Daily" {
			selectedOption = t.DailyDisplay(db)
		} else if selectedOption == "Weekly" {
			selectedOption = t.WeeklyDisplay(db)
		} else if selectedOption == "All Time" {
			selectedOption = t.AllTimeDisplay(db)
		} else {
			selectedOption = t.MenuDisplay()
		}
	}
}
