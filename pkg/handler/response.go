package handler

import (
	"github.com/gin-gonic/gin"
	todo "github.com/prerec/go-final"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Error string `json:"error"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

type getAllTasksResponse struct {
	Tasks []todo.Task `json:"tasks"`
}

type statusResponse struct {
	Status string `json:"status"`
}
