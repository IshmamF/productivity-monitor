package database

import (
	"gorm.io/datatypes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"encoding/json"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("monitor.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}

func AddSessionData(Date datatypes.Date, Start_Time datatypes.Time, End_Time datatypes.Time) {
	db := ConnectDB()
	db.AutoMigrate(&Session{})
	var sessions []Session
	lastID := uint(db.Find(&sessions).RowsAffected)
	db.Create(&Session{ID: lastID + 1, Date: Date, Start_Time: Start_Time, End_Time: End_Time})
}

func ReadAllSessionData() string {
	db := ConnectDB()
	db.AutoMigrate(&Session{})
	var sessions []Session
	b, err := json.Marshal(sessions)
    if err != nil {
        panic("Json Marshal Failed")
    }
    return string(b)
}