package config

import (
	driver "github.com/navisot/go-exchanger/database"
	"log"
)

// ConnectToDatabase creates a new database connection
func ConnectToDatabase(dsn string) (*driver.DBDriver, error) {

	newDB, err := driver.NewDatabase(dsn)

	if err != nil {
		panic(err)
	}

	if err != nil {
		return nil, err
	}

	err = newDB.AutoMigrate()

	if err != nil {
		log.Fatal(err.Error())
	}

	return &driver.DBDriver{Gorm: newDB}, nil
}
