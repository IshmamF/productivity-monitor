package main

import (
	"fmt"
	// "github.com/IshmamF/productivity-monitor/window" Uncomment when using module
	"time"
	"github.com/IshmamF/productivity-monitor/darwin"
	"github.com/IshmamF/productivity-monitor/utils"
	"github.com/IshmamF/productivity-monitor/database"
	"github.com/andybrewer/mack"
)

// TO DO NEXT: 
// [x] Create function for currentTimestamp in Utils
// [x] Look up how to check if table and file exists or not (to handle new and previous users) 
// [x] Create function to process string recieved from GetForegroundWindowData()
// - Get user input on when to send alert, might need to use a counter to keep track of time passed  
// - Need to look into how to execute other functions like viewing statistics or running the alert
// while the logging occurs 
// - Create GUI 
// - Option to see current session data 
// - Convert data to daily/weekly/monthly/all time statistics 

func main() {
	os := utils.Get_OS()
	db := &database.DB{}
	db.Connection()
	startTime := utils.GetCurrentTimestamp()
	var alert_interval int

	fmt.Println("Set Alert Interval (in seconds): ") 
	fmt.Scanln(&alert_interval) 
	
	counter := 0
  	for {
		var activity database.Activity
		var currWindow string
		activity.Start_Time = startTime
		activity.Log_Time = utils.GetCurrentTimestamp()

		if counter > 0 && counter % alert_interval == 0 {
			result := db.HighestUsedApp(activity.Log_Time - int64(alert_interval), activity.Log_Time)
			fmt.Println(result)
			mack.Alert("Productivity Monitor", result.App_Name + " for " + utils.IntToString(int64(result.Count)) + " seconds","critical")
		}

		if os == "darwin" {
			currWindow = darwin.GetForegroundWindowData()
			activity.Url, activity.App_Name, activity.Title, activity.App_Or_Site = utils.ProcessActivityDetails(currWindow)
			fmt.Println("Start: " + utils.IntToString(startTime) + ` Log: ` + utils.IntToString(activity.Log_Time) + ` Window: `, activity.Url, activity.App_Name, activity.Title, activity.App_Or_Site)
		} else {
			// currWindow = window.GetForegroundWindowData() Uncomment when building for use
		}
		db.AddActivity(activity)
		time.Sleep(1 * time.Second)
		counter += 1
	}

	//results := db.CountAppUsageWithRange(1723256375, 1723400234)

}
