package driver

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBDriver struct {
	Gorm *gorm.DB
}

// NewDatabase creates a new database for the application (gorm in this case)
func NewDatabase(dsn string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil

}
