package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo "github.com/prerec/go-final"
	"time"
)

type TodoTaskSqlite struct {
	db *sqlx.DB
}

func NewTodoTaskSqlite(db *sqlx.DB) *TodoTaskSqlite {
	return &TodoTaskSqlite{db: db}
}

func (r *TodoTaskSqlite) Create(task todo.Task) (int, error) {
	createTaskQuery := fmt.Sprintf("INSERT INTO %s (date, title, comment, repeat) VALUES (?, ?, ?, ?)", schedulerTable)
	res, err := r.db.Exec(createTaskQuery, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	return int(id), err
}

func (r *TodoTaskSqlite) GetAll() ([]todo.Task, error) {
	var tasks []todo.Task

	getTasksQuery := fmt.Sprintf("SELECT * FROM %s ORDER BY date DESC", schedulerTable)
	err := r.db.Select(&tasks, getTasksQuery)

	return tasks, err
}

func (r *TodoTaskSqlite) Search(query string) ([]todo.Task, error) {
	var tasks []todo.Task
	var searchTaskQuery string

	parsedDate, err := time.Parse("02.01.2006", query)
	if err == nil {
		searchDate := parsedDate.Format("20060102")
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

func (r *TodoTaskSqlite) GetByID(id int) (todo.Task, error) {
	var task todo.Task

	getTaskByIdQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", schedulerTable)
	err := r.db.Get(&task, getTaskByIdQuery, id)

	return task, err
}
