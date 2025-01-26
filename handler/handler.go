package handler

import (
	"dev.azure.com/lucasfinetti/Finetti/_git/Gopportunities/config"
	"gorm.io/gorm"
)

var (
	Logger *config.Logger
	Db     *gorm.DB
)

func InitializeHandler() {
	Logger = config.GetLogger("handler")
	Db = config.GetSQLServer()
}
