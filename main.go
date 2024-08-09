package main

import (
	"fmt"
	// "github.com/IshmamF/productivity-monitor/window" Uncomment when using module
	"time"

	"github.com/IshmamF/productivity-monitor/darwin"
	"github.com/IshmamF/productivity-monitor/utils"
	//"gorm.io/datatypes"	
	"database/sql"
	_ "github.com/marcboeker/go-duckdb"
	"strconv"

)

type Activity struct {
	Start_Time int64 // primary key 
	Log_Time int64 // timestamp of when activity was recorded
	App_Name string
}

type App_Count struct {
	App_Name string
	Count int
}

func main() {
	os := utils.Get_OS()

	db, err := sql.Open("duckdb", "./storage.db")
	if err != nil {
		fmt.Print("COULD NOT CONNECT TO DATABASE", err)
	}
	defer db.Close()
	


	db.Exec(`CREATE TABLE Activity (Start_Time INTEGER, Log_Time INTEGER, App_Name VARCHAR)`)
	fmt.Println("Database connected")

	startTime := time.Now().Unix()
	fmt.Println("LOOP START")
  	for i := 0; i < 5; i++ {
		var currWindow string
		if os == "darwin" {
			currWindow = darwin.GetForegroundWindowData()
			logTime := time.Now().Unix()
			fmt.Println("Start: " + strconv.FormatInt(startTime, 10) + ` Log: ` + strconv.FormatInt(logTime, 10) + ` Window: ` + currWindow)
			//db.Exec(`INSERT INTO Activity VALUES (` + strconv.FormatInt(startTime, 10) + `,` + strconv.FormatInt(logTime, 10) + `,'` + currWindow + `')`)
		} else {
			// currWindow = window.GetForegroundWindowData() Uncomment when building for use
		}
		time.Sleep(2 * time.Second)
	}

	rows, err := db.Query(`SELECT App_Name, count(App_Name) FROM Activity WHERE Log_Time BETWEEN 1723219430 AND 1723334010 group by App_Name`)
	if err != nil {
		panic("Query Failed")
	}
	defer rows.Close()
	//results := []Activity{}
	results := []App_Count{}
	for rows.Next() {
		//activity := Activity{}
		app_count := App_Count{}
		err := rows.Scan(&app_count.App_Name, &app_count.Count)
		if err != nil {
			fmt.Println(err)
		}
		results = append(results, app_count)
	}
	fmt.Println(results)



}
