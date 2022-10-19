package database

import (
	"log"

	"github.com/PedroSMarcal/todo/constants"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewConnectionTODO() *gorm.DB {
	return openConnectionTodoDatabase()
}

func openConnectionTodoDatabase() *gorm.DB {

	databaseConnection, err := gorm.Open(sqlite.Open(constants.DatabaseName), &gorm.Config{})
	if err != nil {
		log.Fatalln("Could not connect")
	}

	return databaseConnection
}

func StartMigrations() {

	db := openConnectionTodoDatabase()
	go runTodoMigration(db)
}
