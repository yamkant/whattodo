package models

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// TODO: 다른 DB 연결을 위해 설정 객체 주입 방식으로 수정
func ConnectDatabase() {

	database, err := gorm.Open(sqlite.Open(os.Getenv("DB_NAME")), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Todo{})
	if err != nil {
		return
	}
	err = database.AutoMigrate(&User{})
	if err != nil {
		return
	}

	DB = database
}
