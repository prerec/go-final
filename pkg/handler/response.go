package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/prerec/go-final/pkg/models"
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
	Tasks []models.Task `json:"tasks"`
}

type statusResponse struct {
	Status string `json:"status"`
}
