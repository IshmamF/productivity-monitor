package main

import (
    "fmt"
	"github.com/prod-mon/window"
	"time"
)


func main() {
  for {
		text := window.GetForegroundWindowData()
		fmt.Println("window :", text)
		time.Sleep(2 * time.Second)
	}
}