package http_server

import (
	"TaskManager/internal/usecase"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Handler struct {
	taskInteractor *usecase.TaskInteractor
}

func NewHandler(taskInteractor *usecase.TaskInteractor) *Handler {
	return &Handler{
		taskInteractor: taskInteractor,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	tasks := router.Group("/tasks")
	{
		tasks.POST("/", h.CreateTaskHandler)
		tasks.GET("/", h.GetTasksHandler)
		tasks.PUT("/:id/mark", h.MarkTaskHandler)
		tasks.PUT("/:id", h.UpdateTaskHandler)
		tasks.DELETE("/:id", h.DeleteTaskHandler)
	}

	return router
}

type statusResponse struct {
	Status string `json:"status"`
}

func (h *Handler) CreateTaskHandler(c *gin.Context) {
	var requestBody struct {
		Text string `json:"text"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := h.taskInteractor.AddTask(requestBody.Text)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"Задача успешно создана"})
}

func (h *Handler) GetTasksHandler(c *gin.Context) {
	tasks, err := h.taskInteractor.GetTasks()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	
	if len(tasks) == 0 {
		c.String(http.StatusNotFound, "Список задач пуст" )
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) UpdateTaskHandler(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var requestBody struct {
		Text string `json:"text"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err = h.taskInteractor.UpdateTask(taskID, requestBody.Text)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"Задача успешно обновлена"})
}

func (h *Handler) MarkTaskHandler(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err = h.taskInteractor.MarkTask(taskID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"Задача успешно отмечена как выполненная"})
}

func (h *Handler) DeleteTaskHandler(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err = h.taskInteractor.DeleteTask(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"Задача успешно удалена"})
}