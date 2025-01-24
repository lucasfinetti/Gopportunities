package handler

import (
	"dev.azure.com/lucasfinetti/Finetti/_git/Gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLServer()
}
