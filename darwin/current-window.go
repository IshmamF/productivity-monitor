package darwin

import (
	"os/exec"
	"strings"
)

func GetForegroundWindowData() string {
	script := `
	global frontApp, appTitle, windowTitle
	set windowTitle to "" 
	tell application "System Events"
		set frontApp to name of first application process whose frontmost is true
		set appTitle to title of first application process whose frontmost is true
		tell process frontApp
			set windowTitle to value of attribute "AXTitle" of window 1
		end tell
	end tell
	return {appTitle, windowTitle}`
	cmd := exec.Command("osascript", "-e", script)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "No window detected"
	}
	prettyOutput := strings.Replace(string(output), "\n", "", -1)
	return prettyOutput
}
