package darwin

import (
	"fmt"
	"time"
	"os/exec"
	"strings"
)

func GetForegroundWindowData() string {
	script := `
	tell application "System Events"
		set frontApp to name of first application process whose frontmost is true
		tell process frontApp
			get value of attribute "AXTitle" of window 1
		end tell
	end tell`
	cmd := exec.Command("osascript", "-e", script)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "No window detected"
	prettyOutput := strings.Replace(string(output), "\n", "", -1)
	return prettyOutput
}
