package usecase

import (
	"TaskManager/internal/entity"
)

type TaskStorage interface {
	AddTask(newText string) error
	GetTasks() ([]entity.Task, error)
	UpdateTask(id int, newText string) error
	MarkTask(id int) error
	DeleteTask(id int) error
}

type TaskInteractor struct {
	taskStorage TaskStorage
}

func NewTaskInteractor(taskStorage TaskStorage) *TaskInteractor {
	return &TaskInteractor{taskStorage: taskStorage}
}

func (t *TaskInteractor) AddTask(newText string) error {
	return t.taskStorage.AddTask(newText)
}

func (t *TaskInteractor) GetTasks() ([]entity.Task, error) {
	return t.taskStorage.GetTasks()
}

func (t *TaskInteractor) UpdateTask(id int, newText string) error {
	return t.taskStorage.UpdateTask(id, newText)
}

func (t *TaskInteractor) MarkTask (id int) error {
	return t.taskStorage.MarkTask(id)
}

func (t *TaskInteractor) DeleteTask(id int) error {
	return t.taskStorage.DeleteTask(id)
}