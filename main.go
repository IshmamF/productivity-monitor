package main

import (
	"fmt"
	// "github.com/IshmamF/productivity-monitor/window" Uncomment when using module
	"time"

	"github.com/IshmamF/productivity-monitor/darwin"
	"github.com/IshmamF/productivity-monitor/utils"
	//"gorm.io/datatypes"
)

func main() {
	os := utils.Get_OS()

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
