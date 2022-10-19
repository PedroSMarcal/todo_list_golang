package database

import (
	"log"

	"github.com/PedroSMarcal/todo_list_golang/tools"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnectionTODO() *gorm.DB {
	return openConnectionTodoDatabase()
}

func openConnectionTodoDatabase() *gorm.DB {

	databaseConnection, err := gorm.Open(postgres.Open(tools.CreateDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalln("Could not connect")
	}

	return databaseConnection
}

func StartMigrations() {

	db := openConnectionTodoDatabase()
	runTodoMigration(db)
}
