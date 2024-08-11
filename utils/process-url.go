package utils

import (
	"log"
	"net/url"
	"strings"
)

func ConvertUrlToDomain(website_url string) string {
	urlstr := ParseUrlString(website_url)

	hostname := RemovePrefix(urlstr)

	return RemovePaths(hostname)
}

func ParseUrlString(str string) string {
	url, err := url.Parse(str)
	if err != nil {
		log.Fatal(err)
	}
	return url.String()
}
 
func RemovePaths(str string) string {
	if strings.Contains(str, "/") {
		temp := strings.Split(str, "/")
		return temp[0]
	} 
	return str
}

func RemovePrefix(urlstr string) string {
	var hostname string
	if strings.HasPrefix(urlstr, "https") {
		hostname = strings.TrimPrefix(urlstr, "https://")
	} else if strings.HasPrefix(urlstr, "http") {
		hostname = strings.TrimPrefix(urlstr, "http://")
	} else {
		hostname = urlstr
	}

	if strings.HasPrefix(hostname, "www") {
		hostname = strings.TrimPrefix(hostname, "www.")
	}
	return hostname
}