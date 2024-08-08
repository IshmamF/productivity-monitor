package main

import (
	"fmt"
	// "github.com/IshmamF/productivity-monitor/window" Uncomment when using module
	"context"
	"log"
	"time"

	"github.com/IshmamF/productivity-monitor/darwin"
	"github.com/IshmamF/productivity-monitor/utils"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v5"
)


func main() {
	os := utils.Get_OS()
	urlExample := "postgres://local:ishmam@localhost:5432/productivity"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Println("Unable to connect to database: \n", err)
	} else {
		fmt.Println("Database connected!")
	}
	defer conn.Close(context.Background())

	type session_row struct {
		session_id int
		date pgtype.Date
		start_time pgtype.Timestamp
		end_time pgtype.Timestamp
	}

	rows, err := conn.Query(context.Background(), "SELECT * FROM session")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var rowSlice []session_row
	for rows.Next() {
		var r session_row
		err := rows.Scan(&r.session_id, &r.date, &r.start_time, &r.end_time)
		if err != nil {
			log.Fatal(err)
		}
	rowSlice = append(rowSlice, r)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(rowSlice)


  	for {
		var currWindow string
		if os == "darwin" {
			currWindow = darwin.GetForegroundWindowData()
		} else {
			// currWindow = window.GetForegroundWindowData() Uncomment when building for use
		}
		fmt.Println("window :", currWindow)
		time.Sleep(2 * time.Second)
	}
}
