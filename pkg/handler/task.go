package handler

import (
	"github.com/gin-gonic/gin"
	todo "github.com/prerec/go-final"
	"net/http"
	"strconv"
	"time"
)

const (
	timeLayout = "20060102"
)

func (h *Handler) createTask(c *gin.Context) {
	var input todo.Task

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Title == "" {
		newErrorResponse(c, http.StatusBadRequest, "Title is required")
		return
	}

	if input.Date == "" {
		input.Date = time.Now().Format(timeLayout)
	}

	if err := timeValidate(input.Date); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Date < time.Now().Format(timeLayout) && input.Repeat == "" {
		input.Date = time.Now().Format(timeLayout)
	}

	if input.Date < time.Now().Format(timeLayout) && input.Repeat != "" {
		newDate, err := getNextDate(time.Now(), input.Date, input.Repeat, timeLayout)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		input.Date = newDate
	}

	id, err := h.services.TodoTask.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *Handler) getAllTasks(c *gin.Context) {
	tasks, err := h.services.TodoTask.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(tasks) == 0 {
		tasks = make([]todo.Task, 0)
	}

	c.JSON(http.StatusOK, getAllTasksResponse{
		Tasks: tasks,
	})
}

func (h *Handler) getTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	task, err := h.services.TodoTask.GetByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *Handler) updateTask(c *gin.Context) {

}

func (h *Handler) deleteTask(c *gin.Context) {

}

func (h *Handler) nextDate(c *gin.Context) {
	now, err := time.Parse(timeLayout, c.Query("now"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	date := c.Query("date")
	repeat := c.Query("repeat")

	nextDate, err := getNextDate(now, date, repeat, timeLayout)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, nextDate)
}
