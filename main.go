package main

import (
    "fmt"
	"github.com/prod-mon/windows"
	//"github.com/prod-mon/darwin"
	"time"
)


func main() {
  	for {
		text := windows.GetForegroundWindowData()
		fmt.Println("window :", text)
		time.Sleep(2 * time.Second)
	}
}
