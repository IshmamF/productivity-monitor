package main

import (
	"fmt"
	// "github.com/IshmamF/productivity-monitor/window" Uncomment when using module
	"github.com/IshmamF/productivity-monitor/darwin"
	"github.com/IshmamF/productivity-monitor/utils"
	"github.com/IshmamF/productivity-monitor/database"
	_"os"
	_"bufio"
	_"strings"
)

// TO DO NEXT: 
// [x] Create function for currentTimestamp in Utils
// [x] Look up how to check if table and file exists or not (to handle new and previous users) 
// [x] Create function to process string recieved from GetForegroundWindowData()
// [x] Get user input on when to send alert, might need to use a counter to keep track of time passed  
// - Need to look into how to execute other functions like viewing statistics or running the alert
// while the logging occurs 
	/*
	look into channels
	go routine , if forever loop , main routine waits for program to exit
	in each infinite loop, have switch cases for the programs to communicate with each other
	*/
// [x] Stop multiple instances from being ran 
	/*
	Lock file , same location (.local/share on mac)
	os.OpenFile to open file
	os.Stat to check file exists
	*/
// - Create TUI 
// - Option to see current session data 
// - Convert data to daily/weekly/monthly/all time statistics 
// - Option to be a login program, starts running automatically when you login to computer

var (
	system_type = utils.Get_OS()
	db = &database.DB{}
	user_selection int
)

func main() {
	db.Connection()

	choice := make(chan string)
	if system_type == "darwin" {
		go darwin.Start_Tracking(choice, db)
	} else {
		// currWindow = window.GetFo
	}
	//reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Press 1 to Start Tracking")
		fmt.Scan(&user_selection)
		//user_selection, _ := reader.ReadString('\n')
		//user_selection = strings.TrimSpace(user_selection)
		switch user_selection {
		case 1:
			choice <- "start"
		case 2:
			choice <- "stop"
		}
	}

	/*
	for {
		select {
		case input := <- choice:
			fmt.Println(input)
		}
	}*/

	//results := db.CountAppUsageWithRange(1723256375, 1723400234)

}
