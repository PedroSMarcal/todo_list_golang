package database

import (
	"log"

	"github.com/PedroSMarcal/todo/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnectionTODO() *gorm.DB {
	return openConnectionTodoDatabase()
}

func openConnectionTodoDatabase() *gorm.DB {

	databaseConnection, err := gorm.Open(mysql.Open(constants.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalln("Could not connect")
	}

	return databaseConnection
}

func StartMigrations() {

	db := openConnectionTodoDatabase()
	go runTodoMigration(db)
}
