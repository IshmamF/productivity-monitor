package darwin

import (
	"time"
	"github.com/IshmamF/productivity-monitor/utils"
	"github.com/IshmamF/productivity-monitor/database"
)

var (
	activity database.Activity
	currWindow string
	counter = 0
	alert_interval = 5
	alert_on = true
)

func Start_Tracking(choice chan string, db *database.DB) {
	running := false
	
	var startTime int64
	for {
		select {
			case track := <- choice:
				if track == "start" {
					if !running {
						running = true
						startTime = utils.GetCurrentTimestamp()
					}
				} else if track == "stop" {
					if running {
						running = false
					}
				}
			default:
				if running {
					activity.Start_Time = startTime
					activity.Log_Time = utils.GetCurrentTimestamp()
					currWindow = GetForegroundWindowData()
		
					activity.Url, activity.App_Name, activity.Title, activity.App_Or_Site = utils.ProcessActivityDetails(currWindow)
					//fmt.Println("Start: " + utils.IntToString(startTime) + ` Log: ` + utils.IntToString(activity.Log_Time) + ` Window: `, activity.Url, activity.App_Name, activity.Title, activity.App_Or_Site)
					
					db.AddActivity(activity)
					if alert_on && counter > 0 && counter % alert_interval == 0 {
						result := db.HighestUsedApp(activity.Log_Time - int64(alert_interval), activity.Log_Time)
						AlertMostUsedApp(result)
					}
					counter += 1
					time.Sleep(1 * time.Second)
				}
		}
	}
}