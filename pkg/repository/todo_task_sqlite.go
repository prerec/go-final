package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/prerec/go-final/pkg/models"
	"github.com/prerec/go-final/pkg/utils"

	"github.com/jmoiron/sqlx"
)

type TodoTaskSqlite struct {
	db *sqlx.DB
}

const (
	limit            = 10
	searchTimeLayout = "02.01.2006"
)

func NewTodoTaskSqlite(db *sqlx.DB) *TodoTaskSqlite {
	return &TodoTaskSqlite{db: db}
}

func (r *TodoTaskSqlite) Create(task models.Task) (int, error) {
	createTaskQuery := fmt.Sprintf("INSERT INTO %s (date, title, comment, repeat) VALUES (?, ?, ?, ?)", schedulerTable)
	res, err := r.db.Exec(createTaskQuery, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	return int(id), err
}

func (r *TodoTaskSqlite) GetAll() ([]models.Task, error) {
	var tasks []models.Task

	getTasksQuery := fmt.Sprintf("SELECT * FROM %s ORDER BY date DESC LIMIT %d", schedulerTable, limit)
	err := r.db.Select(&tasks, getTasksQuery)

	return tasks, err
}

func (r *TodoTaskSqlite) Search(query string) ([]models.Task, error) {
	var tasks []models.Task
	var searchTaskQuery string

	parsedDate, err := time.Parse(searchTimeLayout, query)
	if err == nil {
		searchDate := parsedDate.Format(utils.TimeLayout)
		searchTaskQuery = fmt.Sprintf("SELECT * FROM %s WHERE date = ?", schedulerTable)
		err = r.db.Select(&tasks, searchTaskQuery, searchDate)
	} else {
		searchTaskQuery = fmt.Sprintf("SELECT * FROM %s WHERE title LIKE ? OR comment LIKE ?", schedulerTable)
		err = r.db.Select(&tasks, searchTaskQuery, "%"+query+"%", "%"+query+"%")
	}

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TodoTaskSqlite) GetByID(id int) (models.Task, error) {
	var task models.Task

	getTaskByIdQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", schedulerTable)
	err := r.db.Get(&task, getTaskByIdQuery, id)

	return task, err
}

func (r *TodoTaskSqlite) Update(id int, task models.Task) error {

	if err := utils.RepeatValidate(task.Repeat); err != nil {
		return err
	}
	if err := utils.TitleValidate(task.Title); err != nil {
		return err
	}
	if err := utils.DateValidate(task.Date); err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET ", schedulerTable)
	var args []interface{}
	hasUpdates := false

	if task.Date != "" {
		query += "date=?, "
		args = append(args, task.Date)
		hasUpdates = true
	}
	if task.Title != "" {
		query += "title=?, "
		args = append(args, task.Title)
		hasUpdates = true
	}
	if task.Comment != "" {
		query += "comment=?, "
		args = append(args, task.Comment)
		hasUpdates = true
	}
	if task.Repeat != "" {
		query += "repeat=?, "
		args = append(args, task.Repeat)
		hasUpdates = true
	}

	if !hasUpdates {
		return errors.New("no fields to update")
	}

	query = query[:len(query)-2] + " WHERE id=?"
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *TodoTaskSqlite) Delete(id int) error {
	deleteTaskQuery := fmt.Sprintf("DELETE FROM %s WHERE id = ?", schedulerTable)
	_, err := r.db.Exec(deleteTaskQuery, id)
	if err != nil {
		return err
	}
	return err
}
