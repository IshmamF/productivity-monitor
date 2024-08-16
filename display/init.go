package display

type Display struct {
	tracking_options_on []string
	tracking_options_off []string
	menu_options []string
	alert_on_options []string
	alert_off_options [] string
	statistic_options []string 
}

func (t *Display) Init() {
	t.InitMenu()

	t.InitTracking()

	t.InitAlert()

	t.InitStatistics()
}

func (t *Display) InitMenu() {
	t.menu_options = append(t.menu_options, "Track Activity")
	t.menu_options = append(t.menu_options, "Statistics")
	t.menu_options = append(t.menu_options, "Alert Settings")
	t.menu_options = append(t.menu_options, "Quit Program")
}

func (t *Display) InitTracking() {
	t.tracking_options_off = append(t.tracking_options_off, "Start Tracking")
	t.tracking_options_off = append(t.tracking_options_off, "Menu")

	t.tracking_options_on = append(t.tracking_options_on, "Stop Tracking")
	t.tracking_options_on = append(t.tracking_options_on, "View Current Session Data")
	t.tracking_options_on = append(t.tracking_options_on, "Menu")
}

func (t *Display) InitAlert() {
	t.alert_off_options = append(t.alert_off_options, "Turn Alerts On")
	t.alert_off_options = append(t.alert_off_options, "Menu")

	t.alert_on_options = append(t.alert_on_options, "Turn Alerts Off")
	t.alert_on_options = append(t.alert_on_options, "Set Interval")
	t.alert_on_options = append(t.alert_on_options, "Menu")
}

func (t *Display) InitStatistics() {
	t.statistic_options = append(t.statistic_options, "Daily")
	t.statistic_options = append(t.statistic_options, "Weekly")
	t.statistic_options = append(t.statistic_options, "All Time")
	t.statistic_options = append(t.statistic_options, "Menu")
}