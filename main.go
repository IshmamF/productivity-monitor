package main

import (
	"fmt"
	// "github.com/IshmamF/productivity-monitor/window" Uncomment when using module
	"time"
	"github.com/IshmamF/productivity-monitor/darwin"
	"github.com/IshmamF/productivity-monitor/utils"
	"github.com/IshmamF/productivity-monitor/database"
)

// TO DO NEXT: 
// [x] Create function for currentTimestamp in Utils
// [x] Look up how to check if table and file exists or not (to handle new and previous users) 
// [x] Create function to process string recieved from GetForegroundWindowData()
// [x] Get user input on when to send alert, might need to use a counter to keep track of time passed  
// - Need to look into how to execute other functions like viewing statistics or running the alert
// while the logging occurs 
	/*
	look into channels
	go routine , if forever loop , main routine waits for program to exit
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
	activity database.Activity
	currWindow string
	counter = 0
	alert_interval int
	db = &database.DB{}
	system_type = utils.Get_OS()
)

func main() {
	db.Connection()

	alert_interval = utils.GetUserInterval()
	startTime := utils.GetCurrentTimestamp()
	
  	for {
		activity.Start_Time = startTime
		activity.Log_Time = utils.GetCurrentTimestamp()

		if system_type == "darwin" {
			currWindow = darwin.GetForegroundWindowData()
			activity.Url, activity.App_Name, activity.Title, activity.App_Or_Site = utils.ProcessActivityDetails(currWindow)
			fmt.Println("Start: " + utils.IntToString(startTime) + ` Log: ` + utils.IntToString(activity.Log_Time) + ` Window: `, activity.Url, activity.App_Name, activity.Title, activity.App_Or_Site)
		} else {
			// currWindow = window.GetForegroundWindowData() Uncomment when building for use
		}
		db.AddActivity(activity)
		if counter > 0 && counter % alert_interval == 0 {
			result := db.HighestUsedApp(activity.Log_Time - int64(alert_interval), activity.Log_Time)
			utils.AlertMostUsedApp(result)
		}
		time.Sleep(1 * time.Second)
		counter += 1
	}

	//results := db.CountAppUsageWithRange(1723256375, 1723400234)

}
