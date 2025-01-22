package config

import (
	"fmt"

	"github.com/arthur404dev/gopportunities/schemas"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	server   = "gopportunities-server.database.windows.net"
	port     = 1433
	user     = "CloudSAbe248903"
	password = "Lf@43767166lflf"
	database = "gopportunities-db"
)

func InitializeSQLServer() (*gorm.DB, error) {

	// Build the connection string
	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", user, password, server, port, database)

	// Create DB and connect
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error opening database connection: %v", err)
		return nil, err
	}

	// Migrate the Schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlserver automigration error: %v", err)
		return nil, err
	}

	// Return the DB
	return db, nil
}
