package main

import (
	"fmt"
	"time"
	"runtime"
	"os/exec"
	"strings"
)

func main() {
	runtime.LockOSThread()

	script := `
    tell application "System Events"
        set frontApp to name of first application process whose frontmost is true
        tell process frontApp
            get value of attribute "AXTitle" of window 1
        end tell
    end tell`
	for {
		cmd := exec.Command("osascript", "-e", script)
		output, _ := cmd.CombinedOutput()
		prettyOutput := strings.Replace(string(output), "\n", "", -1)
		fmt.Println(prettyOutput)
		time.Sleep(2 * time.Second)
	}
}
