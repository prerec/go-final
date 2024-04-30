package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/prerec/go-final/pkg/service"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/nextdate", h.nextDate)

		task := api.Group("/task")
		{
			task.POST("/", h.createTask)
			task.GET("/", h.getTaskByID)
			task.PUT("/", h.updateTask)
			task.POST("/done", h.doneTask)
			task.DELETE("/", h.deleteTask)
		}

		tasks := api.Group("/tasks")
		{
			tasks.GET("/", h.getAllTasks)
		}
	}

	router.Use(func(c *gin.Context) {
		fs := http.FileServer(http.Dir("./web"))
		fs.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	})

	return router
}
