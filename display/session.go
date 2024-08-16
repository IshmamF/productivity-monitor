package display

import(
	"github.com/IshmamF/productivity-monitor/database"
	"github.com/IshmamF/productivity-monitor/utils"
	"github.com/pterm/pterm"
	"fmt"
)

func (t *Display) SessionDisplay(db *database.DB, startTime *int64) (string) {
	currentTime := utils.GetCurrentTimestamp()
	data := db.CountAppUsageWithRange(*startTime, currentTime)
	tablerow := ConvertToTableList(data)

	ShowSessionTable(tablerow)

	var input string
	ShowOptions()
	fmt.Scan(&input)

	return HandleSessionInput(input)
}

func ConvertToTableList(data []database.App_Count) (tablerow [][]string) {
	tablerow = append(tablerow, []string{"App or Site Name", "Usage"})
	for _, activity := range data {
		tablerow = append(tablerow, []string{activity.App_Name, fmt.Sprint(activity.Count)})
	}
	return
}

func ShowSessionTable(tablerow [][]string) {
	pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tablerow).Render()
}

func HandleSessionInput(input string) (selectedOption string) {
	switch input {
	case "p":
		selectedOption = "Track Activity"
	case "m": 
		selectedOption = "Menu"
	default:
		selectedOption = "View Current Session Data"
	}
	return
}