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
)

func Start_Tracking(choice chan string, db *database.DB, startTime *int64, running *bool, ticker *time.Ticker) {
	
	for {
		select {
			case track := <- choice:
				if track == "start" {
					if !*running {
						*running = true
						*startTime = utils.GetCurrentTimestamp()
					}
				} else if track == "stop" {
					if *running {
						counter = 0
						*running = false
					}
				}
			default:
				if *running {
					alert_interval := db.GetAlertSettings()
					activity.Start_Time = *startTime
					activity.Log_Time = utils.GetCurrentTimestamp()
					currWindow = GetForegroundWindowData()
		
					activity.Url, activity.App_Name, activity.Title, activity.App_Or_Site = utils.ProcessActivityDetails(currWindow)
					//fmt.Println("Start: " + utils.IntToString(startTime) + ` Log: ` + utils.IntToString(activity.Log_Time) + ` Window: `, activity.Url, activity.App_Name, activity.Title, activity.App_Or_Site)
					
					db.AddActivity(activity)
					if alert_interval.Alert_On && counter > 0 && counter % alert_interval.Interval == 0 {
						currentTime := utils.GetCurrentTimestamp()
						result := db.HighestUsedApp(activity.Log_Time - int64(alert_interval.Interval), currentTime)
						AlertMostUsedApp(result)
					}
					counter += 1
					<-ticker.C
				}
		}
	}
}