package database

import (
	"github.com/PedroSMarcal/todo/models"
	"gorm.io/gorm"
)

func runTodoMigration(db *gorm.DB) {
	db.AutoMigrate(&models.Task{})
}
