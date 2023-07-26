package app

import (
	"errors"
	"os"
	"reflect"
	"testing"
	ent "TaskManager/internal/entity"
)

var ErrEmptyList = errors.New("Список задач пуст")
var ErrIdNotFound = errors.New("Задача с заданным ID не найдена")
var ErrTaskExist = errors.New("Задача уже существует")

func TestCreateTask(t *testing.T) {
	tm := NewTaskManager()
	tm.Tasks = []ent.Task{
		ent.NewTask(1, "Task 1", false),
		ent.NewTask(2, "Task 2", false),
	}

	expectedTasks := append(tm.Tasks, ent.NewTask(3, "New Task", false))

	input := "New Task\n"
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write([]byte(input))
	w.Close()

	tm.CreateTask()

	if !reflect.DeepEqual(tm.Tasks, expectedTasks) {
		t.Errorf("Expected tasks: %v, but got: %v", expectedTasks, tm.Tasks)
	}
}

func TestErrCreateTask(t *testing.T) {
	tm := NewTaskManager()
	tm.Tasks = []ent.Task{
		ent.NewTask(1, "Task 1", false),
		ent.NewTask(2, "Task 2", false),
	}

	input := "Task 1\n"
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write([]byte(input))
	w.Close()

	err := tm.CreateTask()

	if !reflect.DeepEqual(err, ErrTaskExist) {
		t.Errorf("Expected error: %v, but got: %v", ErrTaskExist, err)
	}
}

func TestUpdateTask(t *testing.T) {
	tm := NewTaskManager()
	tm.Tasks = []ent.Task{
		ent.NewTask(1, "Task 1", false),
		ent.NewTask(2, "Task 2", false),
	}

	expectedTasks := []ent.Task{
		ent.NewTask(1, "Task 1", false),
		ent.NewTask(2, "Updated Task", false),
	}

	input := "2\nUpdated Task\n"
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write([]byte(input))
	w.Close()

	tm.UpdateTask()

	if !reflect.DeepEqual(tm.Tasks, expectedTasks) {
		t.Errorf("Expected tasks: %v, but got: %v", expectedTasks, tm.Tasks)
	}
}

func TestErrUpdateTask1(t *testing.T) {
	tm := NewTaskManager()

	err := tm.UpdateTask()

	if !reflect.DeepEqual(err, ErrEmptyList) {
		t.Errorf("Expcted error: %v, but got: %v", ErrEmptyList, err)
	}
}

func TestErrUpdateTask2(t *testing.T) {
	tm := NewTaskManager()
	tm.Tasks = []ent.Task{
		ent.NewTask(1, "Task 1", false),
		ent.NewTask(2, "Task 2", false),
	}

	input := "3\n"
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write([]byte(input))
	w.Close()

	err := tm.UpdateTask()

	if !reflect.DeepEqual(err, ErrIdNotFound) {
		t.Errorf("Expcted error: %v, but got: %v", ErrIdNotFound, err)
	}
}

func TestMarkTask(t *testing.T) {
	tm := NewTaskManager()
	tm.Tasks = []ent.Task{
		ent.NewTask(1, "Task 1", false),
		ent.NewTask(2, "Task 2", false),
	}

	expectedTasks := []ent.Task{
		ent.NewTask(1, "Task 1", false),
		ent.NewTask(2, "Task 2", true),
	}

	input := "2\n"
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write([]byte(input))
	w.Close()

	tm.MarkTask()

	if !reflect.DeepEqual(tm.Tasks, expectedTasks) {
		t.Errorf("Expected tasks: %v, but got: %v", expectedTasks, tm.Tasks)
	}
}

func TestErrMarkTask1(t *testing.T) {
	tm := NewTaskManager()

	err := tm.UpdateTask()

	if !reflect.DeepEqual(err, ErrEmptyList) {
		t.Errorf("Expected error: %v, but got: %v", ErrEmptyList, err)
	}
}

func TestErrMarkTask2(t *testing.T) {
	tm := NewTaskManager()
	tm.Tasks = []ent.Task{
		ent.NewTask(1, "Task 1", false),
		ent.NewTask(2, "Task 2", false),
	}

	input := "3\n"
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write([]byte(input))
	w.Close()

	err := tm.MarkTask()

	if !reflect.DeepEqual(err, ErrIdNotFound) {
		t.Errorf("Expcted error: %v, but got: %v", ErrIdNotFound, err)
	}
}

func TestDeleteTask(t *testing.T) {
	tm := NewTaskManager()
	tm.Tasks = []ent.Task{
		ent.NewTask(1, "Task 1", false),
		ent.NewTask(2, "Task 2", false),
	}

	expectedTasks := []ent.Task{
		ent.NewTask(1, "Task 1", false),
	}

	input := "2\n"
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write([]byte(input))
	w.Close()

	tm.DeleteTask()

	if !reflect.DeepEqual(tm.Tasks, expectedTasks) {
		t.Errorf("Expected tasks: %v, but got: %v", expectedTasks, tm.Tasks)
	}
}

func TestErrDeleteTask1(t *testing.T) {
	tm := NewTaskManager()

	err := tm.UpdateTask()

	if !reflect.DeepEqual(err, ErrEmptyList) {
		t.Errorf("Expected error: %v, but got: %v", ErrEmptyList, err)
	}
}

func TestErrDeleteTask2(t *testing.T) {
	tm := NewTaskManager()
	tm.Tasks = []ent.Task{
		ent.NewTask(1, "Task 1", false),
		ent.NewTask(2, "Task 2", false),
	}

	input := "3\n"
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write([]byte(input))
	w.Close()

	err := tm.DeleteTask()

	if !reflect.DeepEqual(err, ErrIdNotFound) {
		t.Errorf("Expected error: %v, but got: %v", ErrIdNotFound, err)
	}
}

func TestReadSaveTask(t *testing.T) {
	tm := NewTaskManager()
	tm.Tasks = []ent.Task{
		ent.NewTask(1, "Task 1", false),
		ent.NewTask(2, "Task 2", false),
	}
	
	expectedTasks := []ent.Task{
		ent.NewTask(1, "Task 1", false),
		ent.NewTask(2, "Task 2", false),
	}

	tm.SaveTasks()
	tm.ReadTasks()

	if !reflect.DeepEqual(tm.Tasks, expectedTasks) {
		t.Errorf("Expected tasks: %v, but got: %v", expectedTasks, tm.Tasks)
	}
}
