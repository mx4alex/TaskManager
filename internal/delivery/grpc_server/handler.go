package grpc_server

import (
	"context"
	"TaskManager/api/task"
	"TaskManager/internal/usecase"
)

type Handler struct {
	task.UnimplementedTaskServiceServer

	hostAddr       string
	taskInteractor *usecase.TaskInteractor
}

func NewHandler(hostAddr string, taskInteractor *usecase.TaskInteractor) *Handler {
	return &Handler{
		hostAddr:       hostAddr,
		taskInteractor: taskInteractor,
	}
}

func (h *Handler) CreateTask(ctx context.Context, req *task.CreateTaskRequest) (*task.Status, error) {
	text := req.GetText()

	err := h.taskInteractor.AddTask(text)
	if err != nil {
		return nil, err
	}

	response := &task.Status{
		Message: "Задача успешно создана",
	}

	return response, nil
}

func (h *Handler) GetTasks(req *task.Empty, stream task.TaskService_GetTasksServer) error {
	tasks, err := h.taskInteractor.GetTasks()
	if err != nil {
		return err
	}

	for _, val := range tasks {
		response := &task.Task{
			Id:   int32(val.ID),
			Text: val.Text,
			Done: val.Done,
		}

		if err := stream.Send(response); err != nil {
			return err
		}
	}

	return nil
}

func (h *Handler) UpdateTask(ctx context.Context, req *task.UpdateTaskRequest) (*task.Status, error) {
	taskID := req.GetId()
	newText := req.GetText()

	err := h.taskInteractor.UpdateTask(int(taskID), newText)
	if err != nil {
		return nil, err
	}

	response := &task.Status{
		Message: "Задача успешно обновлена",
	}

	return response, nil
}

func (h *Handler) MarkTask(ctx context.Context, req *task.MarkTaskRequest) (*task.Status, error) {
	taskID := req.GetId()

	err := h.taskInteractor.MarkTask(int(taskID))
	if err != nil {
		return nil, err
	}

	response := &task.Status{
		Message: "Задача успешно отмечена как выполненная",
	}

	return response, nil
}

func (h *Handler) DeleteTask(ctx context.Context, req *task.DeleteTaskRequest) (*task.Status, error) {
	taskID := req.GetId()

	err := h.taskInteractor.DeleteTask(int(taskID))
	if err != nil {
		return nil, err
	}

	response := &task.Status{
		Message: "Задача успешно удалена",
	}

	return response, nil
}