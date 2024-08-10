package main

import (
	"fmt"
	// "github.com/IshmamF/productivity-monitor/window" Uncomment when using module
	"time"

	"github.com/IshmamF/productivity-monitor/darwin"
	"github.com/IshmamF/productivity-monitor/utils"
	"github.com/IshmamF/productivity-monitor/database"
	_ "github.com/marcboeker/go-duckdb"
	"strconv"

)

func main() {
	os := utils.Get_OS()
	db := &database.DB{}
	db.Connection()
	startTime := time.Now().Unix()
	fmt.Println("LOOP START")
  	for i := 0; i < 5; i++ {
		var currWindow string
		logTime := time.Now().Unix()
		if os == "darwin" {
			currWindow = darwin.GetForegroundWindowData()
			fmt.Println("Start: " + strconv.FormatInt(startTime, 10) + ` Log: ` + strconv.FormatInt(logTime, 10) + ` Window: ` + currWindow)
		} else {
			// currWindow = window.GetForegroundWindowData() Uncomment when building for use
		}
		db.AddActivity(startTime, logTime, currWindow)
		time.Sleep(2 * time.Second)
	}

	results := db.CountAppUsageWithRange(1723256375, 1723341501)
	fmt.Println(results)



}
