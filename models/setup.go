package models

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	if os.Getenv("DSN") != "" {
		dsn := os.Getenv("DSN")
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}

		err = db.AutoMigrate(Project{}, Language{}, User{}, Group{}, Role{}, Policy{})
		if err != nil {
			return
		}
	}
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(Project{}, Language{}, User{}, Group{}, Role{}, Policy{})
	if err != nil {
		return
	}

	DB = db
}
