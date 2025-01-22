package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize SQLServer
	db, err = InitializeSQLServer()

	if err != nil {
		return fmt.Errorf("error initializing sqlserver: %v", err)
	}

	return nil
}

func GetSQLServer() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}
