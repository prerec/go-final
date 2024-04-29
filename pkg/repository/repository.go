package repository

import (
	"github.com/jmoiron/sqlx"
	todo "github.com/prerec/go-final"
)

type TodoTask interface {
	Create(task todo.Task) (int, error)
	GetAll() ([]todo.Task, error)
	GetByID(id int) (todo.Task, error)
}

type Repository struct {
	TodoTask
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoTask: NewTodoTaskSqlite(db),
	}
}
