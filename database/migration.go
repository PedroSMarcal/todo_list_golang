package database

import (
	"github.com/PedroSMarcal/todo_list_golang/models"
	"gorm.io/gorm"
)

func runTodoMigration(db *gorm.DB) {
	db.AutoMigrate(&models.Task{})
}
