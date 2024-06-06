package service

import (
	"github.com/prerec/go-final/pkg/models"
	"github.com/prerec/go-final/pkg/repository"
)

type TodoTask interface {
	Create(task models.Task) (int, error)
	GetAll() ([]models.Task, error)
	Search(query string) ([]models.Task, error)
	GetByID(id int) (models.Task, error)
	Update(id int, task models.Task) error
	Delete(id int) error
}

type Service struct {
	TodoTask
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TodoTask: NewTodoTaskService(repos.TodoTask),
	}
}
