package utils

import (
	"strconv"
	"time"
)

// gets timestamp of when function is called (aka current time)
func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

// converts timestamp integer into a string
func IntToString(integer int64) string {
	return strconv.FormatInt(integer, 10)
}