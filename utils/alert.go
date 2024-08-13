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

func GetUserInterval() (interval int) {
	fmt.Println("Set Alert Interval (in seconds): ")
	fmt.Scan(&interval)
	return
}