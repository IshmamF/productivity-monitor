package utils

import (
	"github.com/andybrewer/mack"
	"github.com/IshmamF/productivity-monitor/database"
	"fmt"
)

func AlertMostUsedApp(result database.App_Count) {
	mack.Alert("Productivity Monitor", result.App_Name + " for " + IntToString(int64(result.Count)) + " seconds","critical")
	fmt.Println(result)
}

func GetUserInterval() (alert_interval int) {
	fmt.Print("Set Alert Interval (in seconds): ") 
	fmt.Scanln(&alert_interval) 
	return 
}