package main

import (
    "fmt"
	//"github.com/IshmamF/productivity-monitor/windows"
	"github.com/IshmamF/productivity-monitor/darwin"
	"time"
)


func main() {
  	for {
		text := darwin.GetForegroundWindowData()
		fmt.Println("window :", text)
		time.Sleep(2 * time.Second)
	}
}
