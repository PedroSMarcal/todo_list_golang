package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewConnectionTODO() *gorm.DB {
	return openConnectionTodoDatabase()
}

func openConnectionTodoDatabase() *gorm.DB {

	databaseConnection, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Could not connect")
	}

	return databaseConnection
}

func StartMigrations() {

	db := openConnectionTodoDatabase()
	go runTodoMigration(db)
}
