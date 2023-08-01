package http_server

import (
	"TaskManager/internal/usecase"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/swaggo/gin-swagger" 
    "github.com/swaggo/files" 
	_ "TaskManager/docs"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
type errorResponse struct {
	Message string `json:"message"`
}
type inputBody struct {
	Text string `json:"name"`
}

// @Summary 	Create task
// @Tags 		tasks
// @Description create task
// @ID 			create-task
// @Accept  	json
// @Produce  	json
// @Param 		input body inputBody true "name task"
// @Success 	200 {integer} integer 1
// @Failure 	400,404 {object} errorResponse
// @Failure 	default {object} errorResponse
// @Router /tasks/ [post]
func (h *Handler) CreateTaskHandler(c *gin.Context) {
	var requestBody struct {
		Text string `json:"name"`
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

// @Summary 	Get tasks
// @Tags 		tasks
// @Description get all tasks
// @ID 			get-task
// @Accept  	json
// @Produce  	json
// @Success 	200 {integer} integer 1
// @Failure 	400,404 {object} errorResponse
// @Failure 	default {object} errorResponse
// @Router /tasks/ [get]
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

// @Summary 	Update task
// @Tags 		tasks
// @Description update task
// @ID 			update-task
// @Accept  	json
// @Produce 	json
// @Param 		id path string true "update task by id"
// @Param 		input body inputBody true "name task"
// @Success 	200 {integer} integer 1
// @Failure 	400,404 {object} errorResponse
// @Failure 	default {object} errorResponse
// @Router /tasks/{id} [put]
func (h *Handler) UpdateTaskHandler(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	var requestBody struct {
		Text string `json:"name"`
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

// @Summary 	Mark task
// @Tags 		tasks
// @Description mark task
// @ID 			mark-task
// @Accept  	json
// @Produce  	json
// @Param 		id path string true "mark task by id"
// @Success 	200 {integer} integer 1
// @Failure 	400,404 {object} errorResponse
// @Failure 	default {object} errorResponse
// @Router /tasks/{id}/mark [put]
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

// @Summary 	Delete task
// @Tags 		tasks
// @Description delete task
// @ID 			delete-task
// @Accept  	json
// @Produce  	json
// @Param 		id path string true "delete task by id"
// @Success 	200 {integer} integer 1
// @Failure 	400,404 {object} errorResponse
// @Failure 	default {object} errorResponse
// @Router /tasks/{id} [delete]
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