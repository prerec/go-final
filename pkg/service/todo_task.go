package service

import (
	todo "github.com/prerec/go-final"
	"github.com/prerec/go-final/pkg/repository"
)

type TodoTaskService struct {
	repo repository.TodoTask
}

func NewTodoTaskService(repo repository.TodoTask) *TodoTaskService {
	return &TodoTaskService{repo: repo}
}

func (s *TodoTaskService) Create(task todo.Task) (int, error) {
	return s.repo.Create(task)
}

func (s *TodoTaskService) GetAll() ([]todo.Task, error) {
	return s.repo.GetAll()
}

func (s *TodoTaskService) Search(query string) ([]todo.Task, error) {
	return s.repo.Search(query)
}

func (s *TodoTaskService) GetByID(id int) (todo.Task, error) {
	return s.repo.GetByID(id)
}
