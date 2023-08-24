package models

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database_ *gorm.DB

func init() {

	// Open database connection
	db, err := gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Links{})
	database_ = db
}

type Links struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	ChannelID int
	Link      string
}
type LinksSsimplified struct {
	ID   int
	Link string
}
