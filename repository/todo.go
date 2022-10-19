package repository

import (
	"github.com/PedroSMarcal/todo/models"
	"gorm.io/gorm"
)

type (
	TodoRepository interface {
		GetAll(tasks *models.Task) error
	}

	todoRepository struct {
		db *gorm.DB
	}
)

func (r *todoRepository) GetAll(tasks *models.Task) error {
	result := r.db.Find(&tasks)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
