package darwin

import (
	"os/exec"
	"strings"
)

func GetForegroundWindowData() string {
	/*
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
	return {appTitle, windowTitle}`*/

	test_script := `
	tell application "System Events" to set frontApp to name of first application process whose frontmost is true
	tell application "System Events" to set appTitle to title of first application process whose frontmost is true

	global currentTabUrl, currentTabTitle, windowTitle
	set currentTabUrl to ""
	set currentTabTitle to ""
	set windowTitle to ""

	if (frontApp = "Safari") or (frontApp = "Webkit") then
	  using terms from application "Safari"
		tell application frontApp to set currentTabUrl to URL of front document
		tell application frontApp to set currentTabTitle to name of front document
	  end using terms from
	else if (frontApp = "Google Chrome") or (frontApp = "Google Chrome Canary") or (frontApp = "Chromium") then
	  using terms from application "Google Chrome"
		tell application frontApp to set currentTabUrl to URL of active tab of front window
		tell application frontApp to set currentTabTitle to title of active tab of front window
	  end using terms from
	else if (frontApp = "Arc") then
	  using terms from application "Arc"
		tell application frontApp to set currentTabUrl to URL of active tab of front window
		tell application frontApp to set currentTabTitle to title of active tab of front window
	  end using terms from
	else
		tell application "System Events"
			tell process frontApp
				set windowTitle to value of attribute "AXTitle" of window 1
			end tell
		end tell
	end if
	
	return {currentTabUrl, currentTabTitle, appTitle, windowTitle}
	`
	cmd := exec.Command("osascript", "-e", test_script)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "No window detected"
	}
	prettyOutput := strings.Replace(string(output), "\n", "", -1)
	return prettyOutput
}
