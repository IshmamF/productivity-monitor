package utils

import (
	"strings"
)

var (
	Url = " "
	App_Name = " "
	Title = " "
	App_Or_Site = " "
 )
// website_url, website_title, app_title, window_title
/*
	App_Or_Site string // app_title or website domain
	Url string // website_url
	App_Name string // app_title
	Title string // website_title or windowTitle
*/
func ProcessActivityDetails(activity string) (string, string, string, string) {
	if activity == "No window detected" {
		return Url, activity, Title, activity
	}
	data := strings.Split(activity, ",")
	Url = data[0]
	App_Name := data[2]

	if data[1] == " " {
		Title = data[3]
		App_Or_Site = data[2]
	} else {
		Title = data[1]
		App_Or_Site = ConvertUrlToDomain(Url)
	}

	return Url, App_Name, Title, App_Or_Site

}