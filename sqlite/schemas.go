package database

import (
	"gorm.io/gorm"
	"gorm.io/datatypes"
)	

type Session struct {
	gorm.Model
	ID  uint 
	Date datatypes.Date
	Start_Time datatypes.Time
	End_Time datatypes.Time
}

type Activity struct {
	Start_Time datatypes.Time
	Time_Elapsed string
	App_Name string
}
