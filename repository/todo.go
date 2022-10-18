package repository

type (
	todoRepository struct {
	}

	TodoRepository interface {
		openConnection()
	}
)

func (r *todoRepository) openConnection() {
	r.openConnection()

}
