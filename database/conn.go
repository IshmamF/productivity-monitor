package database

import (
	"database/sql"
	_ "github.com/marcboeker/go-duckdb"
	"log"
)

var (
	filepath = "storage.db"
)

type DB struct {
	conn *sql.DB
}

func (d *DB) Connection () {
	db, err := sql.Open("duckdb", filepath)
	if err != nil {
		log.Panic("COULD NOT CONNECT TO DATABASE: ", err)
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