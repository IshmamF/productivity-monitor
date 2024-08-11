package database

import (
	"log"
	"strconv"
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
	rows, err := db.Query(`SELECT App_Name, count(App_Name) FROM Activity WHERE Log_Time BETWEEN ` + strconv.FormatInt(startTime, 10) + ` AND ` + strconv.FormatInt(endTime, 10) + ` group by App_Name`)
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

// Executes a query to get back all rows within Activity Table
func (d *DB) ReadAllRows() []Activity {
	db := d.conn
	rows, err := db.Query(`SELECT App_Name, Start_Time, Log_Time FROM Activity`)
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
		err := rows.Scan(&activity.App_Name, &activity.Start_Time, &activity.Log_Time)
		if err != nil {
			log.Panic("SCAN FAILED: ", err)
		}
		results = append(results, activity)
	}
	return results
}