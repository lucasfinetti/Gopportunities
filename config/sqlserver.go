package config

import (
	"fmt"

	"dev.azure.com/lucasfinetti/Finetti/_git/Gopportunities/schemas"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	server   = "gopportunities-server.database.windows.net"
	port     = 1433
	user     = "CloudSAbe248903"
	password = "Lf@43767166lflf"
	database = "gopportunities_db"
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
	models := []interface{}{
		&schemas.Opening{},
		&schemas.Candidate{},
	}

	err = db.AutoMigrate(models...)
	if err != nil {
		logger.Errorf("sqlserver automigration error: %v", err)
		return nil, err
	}

	// Return the DB
	return db, nil
}
