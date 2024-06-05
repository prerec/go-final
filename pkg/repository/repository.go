package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/prerec/go-final/pkg/models"
)

type TodoTask interface {
	Create(task models.Task) (int, error)
	GetAll() ([]models.Task, error)
	Search(query string) ([]models.Task, error)
	GetByID(id int) (models.Task, error)
	Update(id int, task models.Task) error
	Delete(id int) error
}

type Repository struct {
	TodoTask
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoTask: NewTodoTaskSqlite(db),
	}
}
