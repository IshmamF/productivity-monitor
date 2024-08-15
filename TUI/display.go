package tui

import (
	"github.com/pterm/pterm"
	"github.com/IshmamF/productivity-monitor/database"
)

type TUI struct {
	tracking_options []string
	menu_options []string
	alert_options []string
}

func (t *TUI) Init() {
	t.menu_options = append(t.menu_options, "Track Activity")
	t.menu_options = append(t.menu_options, "Statistics")
	t.menu_options = append(t.menu_options, "Alert Settings")
	t.menu_options = append(t.menu_options, "Quit Program")

	t.tracking_options = append(t.tracking_options, "Start Tracking")
	t.tracking_options = append(t.tracking_options, "Stop Tracking")
	t.tracking_options = append(t.tracking_options, "View Session Data")
	t.tracking_options = append(t.tracking_options, "Menu")

	t.alert_options = append(t.alert_options, "Turn Alert On")
	t.alert_options = append(t.alert_options, "Turn Alert Off")
}

func (t *TUI) MenuDisplay() (selectedOption string) {
	selectedOption, _ = pterm.DefaultInteractiveSelect.WithOptions(t.menu_options).Show()
	return
}

func (t *TUI) TrackingDisplay(choice chan string) (selectedOption string) {
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

func (t *TUI) AlertSettingsDisplay(db *database.DB) (selectedOption string) {
	selectedOption = ""
	return 
}