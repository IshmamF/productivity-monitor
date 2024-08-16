package display

import (
	"github.com/IshmamF/productivity-monitor/database"
	"github.com/pterm/pterm"
)

func (t *Display) MenuDisplay() (selectedOption string) {
	selectedOption, _ = pterm.DefaultInteractiveSelect.WithOptions(t.menu_options).Show()
	return
}

func (t *Display) TrackingDisplay(choice chan string, running *bool) (selectedOption string) {
	if *running {
		selectedOption, _ = pterm.DefaultInteractiveSelect.WithOptions(t.tracking_options_on).Show()
	} else {
		selectedOption, _ = pterm.DefaultInteractiveSelect.WithOptions(t.tracking_options_off).Show()
	}

	switch selectedOption {
	case "Start Tracking":
		choice <- "start"
	case "Stop Tracking":
		choice <- "stop"
	}
	return
}

func (t *Display) AlertSettingsDisplay(db *database.DB) (selectedOption string) {
	alert_settings := db.GetAlertSettings()
	options := t.ChooseAlertOptions(alert_settings)

	selectedOption, _ = pterm.DefaultInteractiveSelect.WithOptions(options).Show()

	if selectedOption == "Turn Alerts Off" {
		db.UpdateAlertOn(false)
	} else if selectedOption == "Turn Alerts On" {
		db.UpdateAlertOn(true)
	}
	return 
}

func (t *Display) ChooseAlertOptions(alert_settings database.Alert_Settings) (options []string) {
	if alert_settings.Alert_On {
		options = t.alert_on_options
	} else {
		options = t.alert_off_options
	}
	return
}

func (t *Display) StatisticsDisplay(db *database.DB) (selectedOption string) {
	selectedOption, _ = pterm.DefaultInteractiveSelect.WithOptions(t.statistic_options).Show()
	return 
}

func ShowOptions() {
	pterm.Info.WithMessageStyle(pterm.NewStyle(pterm.FgLightMagenta)).Println("Enter 'p' to Return to Previous Page")
	pterm.Info.WithMessageStyle(pterm.NewStyle(pterm.FgLightMagenta)).Println("Enter 'm' to Return to go to Menu Page")
	pterm.Info.WithMessageStyle(pterm.NewStyle(pterm.FgLightMagenta)).Println("Enter any character to Refresh")
	textprinter := pterm.BasicTextPrinter{
		Style: pterm.NewStyle(pterm.FgLightGreen),
	}
	textprinter.Print("Input character: ")
}