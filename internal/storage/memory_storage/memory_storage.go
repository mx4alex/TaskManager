package memory_storage

import (
	"errors"
	"TaskManager/internal/entity"
)

type MemoryTaskStorage struct {
	data []entity.Task
}

func New() (*MemoryTaskStorage, error) {
	return &MemoryTaskStorage{}, nil
}

func (ms *MemoryTaskStorage) AddTask(newText string) error {
	ms.data = append(ms.data, entity.NewTask(len(ms.data) + 1, newText, false))
	return nil
}

func (ms *MemoryTaskStorage) GetTasks() ([]entity.Task, error) {
	return ms.data, nil
}

func (ms *MemoryTaskStorage) UpdateTask(id int, newText string) error {
	var found bool
	
	for i, t := range ms.data {
		if t.ID == id {
			found = true
			ms.data[i].Text = newText
			break
		}
	}
	if !found {
		return errors.New("task not found")
	}

	return nil
}

func (ms *MemoryTaskStorage) MarkTask(id int) error {
	var found bool
	
	for i, t := range ms.data {
		if t.ID == id {
			found = true
			ms.data[i].Done = true
			break
		}
	}
	if !found {
		return errors.New("task not found")
	}

	return nil
}

func (ms *MemoryTaskStorage) DeleteTask(id int) error {
	var found bool
	
	for i, t := range ms.data {
		if t.ID == id {
			found = true
			ms.data = append(ms.data[:i], ms.data[i+1:]...)
			break
		}
	}
	if !found {
		return errors.New("task not found")
	}

	return nil
}
