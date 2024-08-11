package database 

import (
	"log"
	"strings"
)

// Logs computer usage into Activity Table of database 
func (d *DB) AddActivity (activity Activity) {
	db := d.conn
	activity.App_Or_Site = strings.TrimSpace(activity.App_Or_Site)
	activity.App_Name = strings.TrimSpace(activity.App_Name)
	activity.Url = strings.TrimSpace(activity.Url)
	activity.Title = strings.TrimSpace(activity.Title)

	_, err := db.Exec(`INSERT INTO Activity VALUES (?, ?, ?, ?, ?, ?)`, activity.Start_Time, activity.Log_Time, activity.App_Or_Site, activity.App_Name, activity.Url, activity.Title)
	if err != nil {
		log.Panic("INSERT ACTIVITY FAILURE: ", err)
	}
}