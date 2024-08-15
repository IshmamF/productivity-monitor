package database

// every second, time series log of current application being used 
// needed for alerts and gathering data 
type Activity struct {
	Start_Time int64 // timestamp of when session started 
	Log_Time int64 // timestamp of when activity was recorded
	App_Or_Site string // app_title or website domain
	Url string // website_url
	App_Name string // app_title
	Title string // website_title or windowTitle
}

// How often app was used 
type App_Count struct {
	App_Name string
	Count int
}

type Alert_Settings struct {
	Alert_ID int
	Alert_On bool
	Interval int
}