package display

import (
	"strconv"
	"strings"
	"github.com/IshmamF/productivity-monitor/database"
	"github.com/pterm/pterm"
)

func (t *Display) IntervalDisplay(db *database.DB, selectedOption string) (string) {
	alert_settings := db.GetAlertSettings()

	ShowIntervalErrors(selectedOption)
	ShowCurrentMonitorInterval(alert_settings)

	resultStr := GetIntervalInput()

	return HandleIntervalInput(resultStr, db)

}

func HandleIntervalInput(resultStr string, db *database.DB) string {
	resultInt, err := strconv.Atoi(resultStr)
	if err != nil {
		checkErrorString := err.Error()
		if strings.Contains(checkErrorString, "invalid syntax") {
			return "Interval Error(1)"
		} 
    }
	
	if resultInt != -1 && resultInt < 60 {
		return "Interval Error(2)"
	}
	if resultInt == -1 {
		return "Menu"
	}

	db.UpdateAlertInterval(resultInt)
	return "Menu"
}


func ShowCurrentMonitorInterval(alert_settings database.Alert_Settings) {
	if alert_settings.Interval == 300 {
		pterm.Info.Println("Current Monitor Interval: ", alert_settings.Interval, " (Default)")
	} else {
		pterm.Info.Println("Current Monitor Interval: ", alert_settings.Interval)
	}
}

func ShowIntervalErrors(selectedOption string) {
	if selectedOption == "Interval Error(1)" {
		pterm.Warning.Println("Invalid Interval: Not an Integer")
	} else if selectedOption == "Interval Error(2)" {
		pterm.Warning.Println("Invalid Interval: Interval Must be Minimum 60 Seconds")
	}
}

func GetIntervalInput() (resultStr string) {
	pterm.Info.Println("Enter -1 to Return to Main Menu")
	TextPrinter := pterm.InteractiveTextInputPrinter{
		TextStyle: pterm.NewStyle(pterm.FgLightGreen),
		DefaultText: "Input Interval in Seconds: ",
	}
	resultStr, _ = TextPrinter.Show()
	return

}