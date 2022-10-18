package database

import (
	"log"

	"github.com/PedroSMarcal/todo/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnectionTODO() *gorm.DB {
	return openConnectionTodoDatabase()
}

func openConnectionTodoDatabase() *gorm.DB {

	databaseConnection, err := gorm.Open(postgres.New(postgres.Config{DSN: config.Envs.DSN}), &gorm.Config{})
	if err != nil {
		log.Fatalln("Could not connect")
	}

	return databaseConnection
}

func StartMigrations() {

	db := openConnectionTodoDatabase()
	go runTodoMigration(db)
}
