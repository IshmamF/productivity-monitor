package database

import (
	"database/sql"
	"log"

	_ "github.com/marcboeker/go-duckdb"
	"os"
	"strings"
)

var (
	filepath = GetHomeDir() + "/.local/share/productivity.db" 
	folderPath = GetHomeDir() + "/.local/share/"
)

type DB struct {
	conn *sql.DB
}

func (d *DB) Connection () {
	os.MkdirAll(folderPath, os.ModePerm)
	db, err := sql.Open("duckdb", filepath)
	if err != nil {
		checkErrorString := err.Error()
		if strings.Contains(checkErrorString, "set lock on file") {
			log.Println("FILE LOCKED: ANOTHER INSTANCE IS RUNNING")
			os.Exit(1)
		} else {
			log.Panic(err)
		}
	}
	d.conn = db
	d.CreateTables()
}

func (d *DB) CreateTables () {
	db := d.conn
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS Activity (Start_Time INTEGER, Log_Time INTEGER, App_Or_Site string, App_Name VARCHAR, Url string, Title string)`)
	if err != nil {
		log.Panic("ADDING Activity TABLE FAILURE: ", err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Alert_Settings (Alert_ID INTEGER, Alert_On BOOLEAN, Interval INTEGER)`)
	if err != nil {
		log.Panic("ADDING Alert_Settings TABLE FAILURE: ", err)
	}
	d.InitAlert()
	
}

func (d *DB) InitAlert() {
	db := d.conn 
	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM Alert_Settings")
	_ = row.Scan(&count)
	if count == 0 {
		_, err := db.Exec(`INSERT INTO Alert_Settings VALUES (?, ?, ?)`, 0, false, 300)
		if err != nil {
			log.Panic("INSERT ALERT SETTINGS FAILURE : ", err)
		}
	}
}

func GetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return homeDir
}