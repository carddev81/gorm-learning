package main

import (
	"learn/gormtestrunner/src/database"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Run gorm tester
func main() {
	//string in this format is used for getting db connection
	dsn := "host=localhost user=test password=itsd dbname=test"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Unable to open connection to database.")
	} //end if

	db.Debug().AutoMigrate(database.TheTables...)
} //end func
