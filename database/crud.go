package database

import (
	"log"
	"strconv"
	"database/sql"
	_ "github.com/marcboeker/go-duckdb"
)

func (d *DB) AddActivity (startTime int64, logTime int64, currWindow string) {
	db := d.conn
	_, err := db.Exec(`INSERT INTO Activity VALUES (` + strconv.FormatInt(startTime, 10) + `,` + strconv.FormatInt(logTime, 10) + `,'` + currWindow + `')`)
	if err != nil {
		log.Panic("INSERT ACTIVITY FAILURE: ", err)
	}
}

func (d *DB) CountAppUsageWithRange (startTime int64, endTime int64) []App_Count {
	rows := d.QueryAppUsageCount(startTime, endTime)

	return ScanAppCountQuery(rows)
}

func (d *DB) QueryAppUsageCount (startTime int64, endTime int64) *sql.Rows {
	db := d.conn
	rows, err := db.Query(`SELECT App_Name, count(App_Name) FROM Activity WHERE Log_Time BETWEEN ` + strconv.FormatInt(startTime, 10) + ` AND ` + strconv.FormatInt(endTime, 10) + ` group by App_Name`)
	if err != nil {
		panic("Query Failed")
	}
	return rows
}

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

func (d *DB) ReadAllRows() []Activity {
	db := d.conn
	rows, err := db.Query(`SELECT App_Name, Start_Time, Log_Time FROM Activity`)
	if err != nil {
		panic("Query Failed")
	}
	return ScanAllRows(rows)

}

func ScanAllRows (rows *sql.Rows) []Activity {
	results := []Activity{}
	for rows.Next() {
		activity := Activity{}
		err := rows.Scan(&activity.Start_Time, &activity.Log_Time, &activity.App_Name)
		if err != nil {
			log.Panic("SCAN FAILED: ", err)
		}
		results = append(results, activity)
	}
	return results
}
