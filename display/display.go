package display

import (
	"github.com/IshmamF/productivity-monitor/database"
	"github.com/pterm/pterm"
)

type Display struct {
	tracking_options []string
	menu_options []string
	alert_on_options []string
	alert_off_options [] string
}

func (t *Display) Init() {
	t.menu_options = append(t.menu_options, "Track Activity")
	t.menu_options = append(t.menu_options, "Statistics")
	t.menu_options = append(t.menu_options, "Alert Settings")
	t.menu_options = append(t.menu_options, "Quit Program")

	t.tracking_options = append(t.tracking_options, "Start Tracking")
	t.tracking_options = append(t.tracking_options, "Stop Tracking")
	t.tracking_options = append(t.tracking_options, "View Session Data")
	t.tracking_options = append(t.tracking_options, "Menu")

	t.alert_off_options = append(t.alert_off_options, "Turn Alerts On")
	t.alert_off_options = append(t.alert_off_options, "Menu")

	t.alert_on_options = append(t.alert_on_options, "Turn Alerts Off")
	t.alert_on_options = append(t.alert_on_options, "Set Interval")
	t.alert_on_options = append(t.alert_on_options, "Menu")
}

func (t *Display) MenuDisplay() (selectedOption string) {
	selectedOption, _ = pterm.DefaultInteractiveSelect.WithOptions(t.menu_options).Show()
	return
}

func (t *Display) TrackingDisplay(choice chan string) (selectedOption string) {
	selectedOption, _ = pterm.DefaultInteractiveSelect.WithOptions(t.tracking_options).Show()

	switch selectedOption {
	case "Start Tracking":
		t.tracking_options[0] = "Currently Tracking"
		choice <- "start"
	case "Stop Tracking":
		t.tracking_options[0] = "Start Tracking"
		choice <- "stop"
	}
	return
}

func (t *Display) AlertSettingsDisplay(db *database.DB) (selectedOption string) {
	alert_settings := db.GetAlertSettings()
	var options []string
	if alert_settings.Alert_On {
		options = t.alert_on_options
	} else {
		options = t.alert_off_options
	}
	selectedOption, _ = pterm.DefaultInteractiveSelect.WithOptions(options).Show()
	if selectedOption == "Turn Alerts Off" {
		db.UpdateAlertOn(false)
	} else if selectedOption == "Turn Alerts On" {
		db.UpdateAlertOn(true)
	}
	return 
}

func (t *Display) IntervalDisplay(db *database.DB) (selectedOption string){
	alert_settings := db.GetAlertSettings() 
	pterm.Info.Println("Current Monitor Interval: ", alert_settings.Interval)
	pterm.DefaultInteractiveTextInput.Show()
	selectedOption = "Done"
	return
}