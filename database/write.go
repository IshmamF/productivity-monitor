package database 

import (
	"log"
	"strconv"
)

// Logs computer usage into Activity Table of database 
func (d *DB) AddActivity (startTime int64, logTime int64, currWindow string) {
	db := d.conn
	_, err := db.Exec(`INSERT INTO Activity VALUES (` + strconv.FormatInt(startTime, 10) + `,` + strconv.FormatInt(logTime, 10) + `,'` + currWindow + `')`)
	if err != nil {
		log.Panic("INSERT ACTIVITY FAILURE: ", err)
	}
}