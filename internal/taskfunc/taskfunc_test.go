package taskfunc

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
	tasks := []ent.Task{
		ent.NewTask(1, "Task 1", false),
		ent.NewTask(2, "Task 2", false),
	}

	expectedTasks := append(tasks, ent.NewTask(3, "New Task", false))

	input := "New Task\n"
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write([]byte(input))
	w.Close()

	CreateTask(&tasks)

	if !reflect.DeepEqual(tasks, expectedTasks) {
		t.Errorf("Expected tasks: %v, but got: %v", expectedTasks, tasks)
	}
}

func TestErrCreateTask(t *testing.T) {
	tasks := []ent.Task{
		ent.NewTask(1, "Task 1", false),
		ent.NewTask(2, "Task 2", false),
	}

	input := "Task 1\n"
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write([]byte(input))
	w.Close()

	err := CreateTask(&tasks)

	if !reflect.DeepEqual(err, ErrTaskExist) {
		t.Errorf("Expected error: %v, but got: %v", ErrTaskExist, err)
	}
}

func TestUpdateTask(t *testing.T) {
	tasks := []ent.Task{
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

	UpdateTask(&tasks)

	if !reflect.DeepEqual(tasks, expectedTasks) {
		t.Errorf("Expected tasks: %v, but got: %v", expectedTasks, tasks)
	}
}

func TestErrUpdateTask1(t *testing.T) {
	tasks := []ent.Task{}

	err := UpdateTask(&tasks)

	if !reflect.DeepEqual(err, ErrEmptyList) {
		t.Errorf("Expcted error: %v, but got: %v", ErrEmptyList, err)
	}
}

func TestErrUpdateTask2(t *testing.T) {
	tasks := []ent.Task{
		ent.NewTask(1, "Task 1", false),
		ent.NewTask(2, "Task 2", false),
	}

	input := "3\n"
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write([]byte(input))
	w.Close()

	err := UpdateTask(&tasks)

	if !reflect.DeepEqual(err, ErrIdNotFound) {
		t.Errorf("Expcted error: %v, but got: %v", ErrIdNotFound, err)
	}
}

func TestMarkTask(t *testing.T) {
	tasks := []ent.Task{
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

	MarkTask(&tasks)

	if !reflect.DeepEqual(tasks, expectedTasks) {
		t.Errorf("Expected tasks: %v, but got: %v", expectedTasks, tasks)
	}
}

func TestErrMarkTask1(t *testing.T) {
	tasks := []ent.Task{}

	err :=  MarkTask(&tasks)

	if !reflect.DeepEqual(err, ErrEmptyList) {
		t.Errorf("Expcted error: %v, but got: %v", ErrEmptyList, err)
	}
}

func TestErrMarkTask2(t *testing.T) {
	tasks := []ent.Task{
		ent.NewTask(1, "Task 1", false),
		ent.NewTask(2, "Task 2", false),
	}

	input := "3\n"
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write([]byte(input))
	w.Close()

	err := MarkTask(&tasks)

	if !reflect.DeepEqual(err, ErrIdNotFound) {
		t.Errorf("Expcted error: %v, but got: %v", ErrIdNotFound, err)
	}
}

func TestDeleteTask(t *testing.T) {
	tasks := []ent.Task{
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

	DeleteTask(&tasks)

	if !reflect.DeepEqual(tasks, expectedTasks) {
		t.Errorf("Expected tasks: %v, but got: %v", expectedTasks, tasks)
	}
}

func TestErrDeleteTask1(t *testing.T) {
	tasks := []ent.Task{}

	err := DeleteTask(&tasks)

	if !reflect.DeepEqual(err, ErrEmptyList) {
		t.Errorf("Expcted error: %v, but got: %v", ErrEmptyList, err)
	}
}

func TestErrDeleteTask2(t *testing.T) {
	tasks := []ent.Task{
		ent.NewTask(1, "Task 1", false),
		ent.NewTask(2, "Task 2", false),
	}

	input := "3\n"
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write([]byte(input))
	w.Close()

	err := DeleteTask(&tasks)

	if !reflect.DeepEqual(err, ErrIdNotFound) {
		t.Errorf("Expcted error: %v, but got: %v", ErrIdNotFound, err)
	}
}

func TestReadSaveTask(t *testing.T) {
	expectedTasks := []ent.Task{
		ent.NewTask(1, "Task 1", false),
		ent.NewTask(2, "Task 2", false),
	}

	SaveTasks(expectedTasks)
	tasks, _ := ReadTasks()

	if !reflect.DeepEqual(tasks, expectedTasks) {
		t.Errorf("Expcted tasks: %v, but got: %v", expectedTasks, tasks)
	}
}