package storage


// every second, time series log of current application being used 
// needed for alerts and gathering data 
type Activity struct {
	Start_Time int64 // primary key 
	Log_Time int64 // timestamp of when activity was recorded
	App_Name string
}

// braekdown of apps used within a session
// would be used to view statistics across days/weeks/months
// not sure if i need this
type Time_Elapsed struct {
	Start_Time int64
	Elapsed int64 
	App_Name string
}