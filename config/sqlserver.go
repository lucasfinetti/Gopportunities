package config

import (
	"fmt"

	"dev.azure.com/lucasfinetti/Finetti/_git/Gopportunities/schemas"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	server   = "dpg-cvntfsk9c44c73fg9n10-a.oregon-postgres.render.com"
	port     = 5432
	user     = "gopportunities_db_user"
	password = "8gRZYmsZiEWyHzPMXXsSQYYJy5DOaaIu"
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
