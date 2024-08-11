package database

import (
	"log"
	"database/sql"
)

// Finds the number of occurances of an app/site within a time period
func (d *DB) CountAppUsageWithRange (startTime int64, endTime int64) []App_Count {
	rows := d.QueryAppUsageCount(startTime, endTime)

	return ScanAppCountQuery(rows)
}

// Executes the query to get the count of number of app/site within time period
func (d *DB) QueryAppUsageCount (startTime int64, endTime int64) *sql.Rows {
	db := d.conn
	rows, err := db.Query(`SELECT App_Name, COUNT(App_Name) FROM Activity WHERE Log_Time BETWEEN ? AND ? group by App_Name`, startTime, endTime)
	if err != nil {
		panic("Query Failed")
	}
	return rows
}

// Scans the data returned from query and places into struct
func ScanAppCountQuery (rows *sql.Rows) []App_Count {
	results := []App_Count{}
	for rows.Next() {
		//activity := Activity{}
		app_count := App_Count{}
		err := rows.Scan(&app_count.App_Name, &app_count.Count)
		if err != nil {
			log.Panic("SCAN FAILED: ", err)
		}
		results = append(results, app_count)
	}
	return results
}

func (d *DB) HighestUsedApp(startTime int64, endTime int64) App_Count {
	db := d.conn
	query := `
	SELECT App_Or_Site, COUNT(App_Or_Site)
	FROM Activity 
	WHERE Log_Time BETWEEN ? AND ? 
	GROUP BY App_Or_Site
	ORDER BY COUNT(App_Or_Site) DESC LIMIT 1;
	`
	var appCount App_Count
	row := db.QueryRow(query, startTime, endTime)
	_ = row.Scan(&appCount.App_Name, &appCount.Count)
	return appCount
}

// Executes a query to get back all rows within Activity Table
func (d *DB) ReadAllRows() []Activity {
	db := d.conn
	rows, err := db.Query(`SELECT Title, App_Or_Site, Url, App_Name, Start_Time, Log_Time FROM Activity`)
	if err != nil {
		panic("Query Failed")
	}
	return ScanAllRows(rows)

}

// Scans the data returned from query to get all rows 
func ScanAllRows (rows *sql.Rows) []Activity {
	results := []Activity{}
	for rows.Next() {
		activity := Activity{}
		err := rows.Scan(&activity.Title, &activity.App_Or_Site, &activity.Url, &activity.App_Name, &activity.Start_Time, &activity.Log_Time)
		if err != nil {
			log.Panic("SCAN FAILED: ", err)
		}
		results = append(results, activity)
	}
	return results
}
