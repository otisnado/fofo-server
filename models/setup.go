package models

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbhost := os.Getenv("DBHOST")
	dbport := os.Getenv("DBPORT")
	dbname := os.Getenv("DBNAME")
	dbuser := os.Getenv("DBUSER")
	dbpass := os.Getenv("DBPASS")
	dbtz := os.Getenv("DBTZ")

	dsn := dbuser + ":" + dbpass + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?parseTime=true&loc=" + dbtz
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(Project{})
	if err != nil {
		return
	}

	DB = db
}
