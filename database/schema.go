package database

// every second, time series log of current application being used 
// needed for alerts and gathering data 
type Activity struct {
	Start_Time int64 // timestamp of when session started 
	Log_Time int64 // timestamp of when activity was recorded
	App_Name string // website_url, website_title, app_title, windowTitle
}

// How often app was used 
type App_Count struct {
	App_Name string
	Count int
}