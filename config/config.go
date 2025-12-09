package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error

	db, err = InitializeMySQL()

	if err != nil {
		return fmt.Errorf("error initializing database: %v", err)
	}

	return nil
}
