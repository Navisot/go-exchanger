package database

import (
	"log"
)

// ConnectToDatabase creates a new database connection
func ConnectToDatabase(dsn string) (*DBDriver, error) {

	newDB, err := NewDatabase(dsn)

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

	return &DBDriver{Gorm: newDB}, nil
}
