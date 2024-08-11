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
	fmt.Println("LOOP START")
  	for i := 0; i < 5; i++ {
		var currWindow string
		logTime := utils.GetCurrentTimestamp()
		var activity database.Activity
		if os == "darwin" {
			currWindow = darwin.GetForegroundWindowData()
			activity.Url, activity.App_Name, activity.Title, activity.App_Or_Site = utils.ProcessActivityDetails(currWindow)
			fmt.Println("Start: " + utils.IntToString(startTime) + ` Log: ` + utils.IntToString(logTime) + ` Window: `, activity.Url, activity.App_Name, activity.Title, activity.App_Or_Site)
		} else {
			// currWindow = window.GetForegroundWindowData() Uncomment when building for use
		}
		db.AddActivity(startTime, logTime, currWindow)
		time.Sleep(2 * time.Second)
	}

	results := db.CountAppUsageWithRange(1723256375, 1723341501)
	//results := db.ReadAllRows()
	fmt.Println(results)

}
