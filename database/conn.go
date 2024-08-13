package database

import (
	"database/sql"
	//"fmt"
	"log"

	_ "github.com/marcboeker/go-duckdb"

	//	"io"
	//	"errors"
	"os"
	"strings"
)

var (
	filepath = GetHomeDir() + "/.local/share/productivity.db" 
)

type DB struct {
	conn *sql.DB
}

func (d *DB) Connection () {
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
		log.Panic("ADDING TABLE FAILURE: ", err)
	}
}

func GetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return homeDir
}