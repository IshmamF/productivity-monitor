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
)

// TO DO NEXT: 
// [x] Create function for currentTimestamp in Utils
// [x] Look up how to check if table and file exists or not (to handle new and previous users) 
// [x] Create function to process string recieved from GetForegroundWindowData()
// [x] Get user input on when to send alert, might need to use a counter to keep track of time passed  
// [x] Need to look into how to execute other functions like viewing statistics or running the alert
// while the logging occurs 
	/*
	look into channels
	go routine , if forever loop , main routine waits for program to exit
	in each infinite loop, have switch cases for the programs to communicate with each other
	*/
// [x] Stop multiple instances from being ran 
	/*
	Lock file , same location (.local/share on mac)
	os.OpenFile to open file
	os.Stat to check file exists
	*/
// - Create TUI 
// - Option to see current session data 
// - Convert data to daily/weekly/monthly/all time statistics 
// - Option to be a login program, starts running automatically when you login to computer

var (
	db = &database.DB{}
	t = &display.Display{}
	selectedOption string
)

func main() {
	db.Connection()

	t.Init()

	choice := make(chan string)
	go darwin.Start_Tracking(choice, db)

	for {
		screen.Clear()
		screen.MoveTopLeft()
		pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgBlack)).WithTextStyle(pterm.NewStyle(pterm.FgLightCyan)).Println("Productivity Monitor")
		if strings.Contains(strings.ToLower(selectedOption),"track") {
			selectedOption = t.TrackingDisplay(choice)
		} else if strings.Contains(strings.ToLower(selectedOption),"alert")  {
			selectedOption = t.AlertSettingsDisplay(db)
		} else if strings.Contains(strings.ToLower(selectedOption), "interval") {
			selectedOption = t.IntervalDisplay(db, selectedOption)
		} else {
			selectedOption = t.MenuDisplay()
		}
	}


	//results := db.CountAppUsageWithRange(1723256375, 1723400234)

}
