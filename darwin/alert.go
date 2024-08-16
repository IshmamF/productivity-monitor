package darwin

import (
	"github.com/andybrewer/mack"
	"github.com/IshmamF/productivity-monitor/database"
	"github.com/IshmamF/productivity-monitor/utils"
)

func AlertMostUsedApp(result database.App_Count) {
	mack.Alert("Productivity Monitor", result.App_Name + " for " + utils.IntToString(int64(result.Count)) + " seconds","critical")
}
