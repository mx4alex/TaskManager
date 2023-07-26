package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"errors"
	ent "TaskManager/internal/entity"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const tasksDB = "storage/tasks.db"

type TaskManager struct {
	Tasks []ent.Task
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (tm *TaskManager) CommandPrint() {
	fmt.Println("Выберите действие:")
	fmt.Println("Create - Создать задачу")
	fmt.Println("Read - Показать список задач")
	fmt.Println("Update - Обновить задачу")
	fmt.Println("Mark - Отметить задачу выполненной")
	fmt.Println("Delete - Удалить задачу")
	fmt.Println("Exit - Выйти")
	fmt.Println()
}

func (tm *TaskManager) CreateTask() error {
	fmt.Println("Введите задачу:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	for _, t := range tm.Tasks {
		if t.Text == text {
			return errors.New("Задача уже существует")
		}
	}

	tm.Tasks = append(tm.Tasks, ent.NewTask(len(tm.Tasks) + 1, text, false))

	if err := tm.SaveTasks(); err != nil {
		return errors.New("Ошибка сохранения задачи")
	}

	fmt.Println("Задача успешно создана:", text)

	return nil
}

func (tm *TaskManager) PrintTasks() error {
	if len(tm.Tasks) == 0 {
		return errors.New("Список задач пуст")
	}

	fmt.Println("Список задач:")
	for _, t := range tm.Tasks {
		if t.Done {
			fmt.Printf("%d. [x] %s\n", t.ID, t.Text)
		} else {
			fmt.Printf("%d. [ ] %s\n", t.ID, t.Text)
		}
	}

	return nil
}

func (tm *TaskManager) UpdateTask() error {
	if len(tm.Tasks) == 0 {
		return errors.New("Список задач пуст")
	}

	tm.PrintTasks()

	fmt.Println("Введите ID задачи для обновления:")
	var id int
	fmt.Scanf("%d\n", &id)

	var found bool
	for i, t := range tm.Tasks {
		if t.ID == id {
			found = true
			fmt.Println("Введите новый текст задачи:")
			reader := bufio.NewReader(os.Stdin)
			newText, _ := reader.ReadString('\n')
			newText = strings.TrimSpace(newText)

			tm.Tasks[i].Text = newText
			break
		}
	}

	if !found {
		return errors.New("Задача с заданным ID не найдена")
	}

	if err := tm.SaveTasks(); err != nil {
		return errors.New("Ошибка сохранения задачи")
	}

	fmt.Println("Задача успешно обновлена.")

	return nil
}

func (tm *TaskManager) MarkTask() error {
	if len(tm.Tasks) == 0 {
		return errors.New("Список задач пуст")
	}

	tm.PrintTasks()

	fmt.Println("Введите ID задачи, которую нужно отметить как выполненную:")
	var id int
	fmt.Scanln(&id)

	var found bool
	for i, t := range tm.Tasks {
		if t.ID == id {
			found = true
			tm.Tasks[i].Done = true
			break
		}
	}

	if !found {
		return errors.New("Задача с заданным ID не найдена")
	}

	if err := tm.SaveTasks(); err != nil {
		return errors.New("Ошибка сохранения задачи")
	}

	fmt.Println("Задача успешно отмечена как выполненная.")

	return nil
}

func (tm *TaskManager) DeleteTask() error {
	if len(tm.Tasks) == 0 {
		return errors.New("Список задач пуст")
	}

	tm.PrintTasks()
	fmt.Println("Введите ID задачи для удаления:")

	var id int
	fmt.Scanln(&id)

	var found bool
	for i, t := range tm.Tasks {
		if t.ID == id {
			found = true
			tm.Tasks = append(tm.Tasks[:i], tm.Tasks[i+1:]...)
			break
		}
	}

	if !found {
		return errors.New("Задача с заданным ID не найдена")
	}

	if err := tm.SaveTasks(); err != nil {
		return errors.New("Ошибка сохранения задачи")
	}

	fmt.Println("Задача успешно удалена.")

	return nil
}

func (tm *TaskManager) ReadTasks() error {
	db, err := sql.Open("sqlite3", tasksDB)
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, text, done FROM tasks")
	if err != nil {
		return err
	}
	defer rows.Close()

	tm.Tasks = nil

	for rows.Next() {
		var id int
		var text string
		var done bool

		err := rows.Scan(&id, &text, &done)
		if err != nil {
			return err
		}

		tm.Tasks = append(tm.Tasks, ent.NewTask(id, text, done))
	}

	return nil
}

func (tm *TaskManager) SaveTasks() error {
	db, err := sql.Open("sqlite3", tasksDB)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY, text TEXT, done INTEGER)")
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM tasks")
	if err != nil {
		return err
	}

	statement, err := db.Prepare("INSERT INTO tasks (id, text, done) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	for _, t := range tm.Tasks {
		_, err = statement.Exec(t.ID, t.Text, t.Done)
		if err != nil {
			return err
		}
	}

	return nil
}
