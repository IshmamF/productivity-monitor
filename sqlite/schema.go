package database

import (
	"time"
)	

type Session struct {
	ID  uint 
	Date time.Time
	Start_Time int
	End_Time int
}
