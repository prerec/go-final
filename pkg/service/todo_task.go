package service

import (
	"github.com/prerec/go-final/pkg/models"
	"github.com/prerec/go-final/pkg/repository"
)

type TodoTaskService struct {
	repo repository.TodoTask
}

func NewTodoTaskService(repo repository.TodoTask) *TodoTaskService {
	return &TodoTaskService{repo: repo}
}

func (s *TodoTaskService) Create(task models.Task) (int, error) {
	return s.repo.Create(task)
}

func (s *TodoTaskService) GetAll() ([]models.Task, error) {
	return s.repo.GetAll()
}

func (s *TodoTaskService) Search(query string) ([]models.Task, error) {
	return s.repo.Search(query)
}

func (s *TodoTaskService) GetByID(id int) (models.Task, error) {
	return s.repo.GetByID(id)
}

func (s *TodoTaskService) Update(id int, task models.Task) error {
	return s.repo.Update(id, task)
}

func (s *TodoTaskService) Delete(id int) error {
	return s.repo.Delete(id)
}
