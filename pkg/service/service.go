package service

import (
	todo "github.com/prerec/go-final"
	"github.com/prerec/go-final/pkg/repository"
)

type TodoTask interface {
	Create(task todo.Task) (int, error)
	GetAll() ([]todo.Task, error)
	Search(query string) ([]todo.Task, error)
	GetByID(id int) (todo.Task, error)
}

type Service struct {
	TodoTask
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TodoTask: NewTodoTaskService(repos.TodoTask),
	}
}
