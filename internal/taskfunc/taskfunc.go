package taskfunc

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

func CommandPrint () {
	fmt.Println("Выберите действие:")
	fmt.Println("Create - Создать задачу")
	fmt.Println("Read - Показать список задач")
	fmt.Println("Update - Обновить задачу")
	fmt.Println("Mark - Отметить задачу выполненной")
	fmt.Println("Delete - Удалить задачу")
	fmt.Println("Exit - Выйти")
	fmt.Println()
}

func CreateTask(Tasks *[]ent.Task) error {
	fmt.Println("Введите задачу:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	for _, t := range *Tasks {
		if t.Text == text {
			return errors.New("Задача уже существует")
		}
	}

	*Tasks = append(*Tasks, ent.NewTask(len(*Tasks) + 1, text, false))

	if err := SaveTasks(*Tasks); err != nil {
		return errors.New("Ошибка сохранения задачи")
	}

	fmt.Println("Задача успешно создана:", text)

	return nil
}

func PrintTasks(Tasks *[]ent.Task) error {
	if len(*Tasks) == 0 {
		return errors.New("Список задач пуст")
	}

	fmt.Println("Список задач:")
	for _, t := range *Tasks {
		if t.Done {
			fmt.Printf("%d. [x] %s\n", t.ID, t.Text)
		} else {
			fmt.Printf("%d. [ ] %s\n", t.ID, t.Text)
		}
	}

	return nil
}

func UpdateTask(Tasks *[]ent.Task) error {
	if len(*Tasks) == 0 {
		return errors.New("Список задач пуст")
	}

	PrintTasks(Tasks)

	fmt.Println("Введите ID задачи для обновления:")
	var id int
	fmt.Scanf("%d\n", &id)

	var found bool
	for i, t := range *Tasks {
		if t.ID == id {
			found = true
			fmt.Println("Введите новый текст задачи:")
			reader := bufio.NewReader(os.Stdin)
			newText, _ := reader.ReadString('\n')
			newText = strings.TrimSpace(newText)

			(*Tasks)[i].Text = newText
			break
		}
	}

	if !found {
		return errors.New("Задача с заданным ID не найдена")
	}

	if err := SaveTasks(*Tasks); err != nil {
		return errors.New("Ошибка сохранения задачи")
	}

	fmt.Println("Задача успешно обновлена.")

	return nil
}

func MarkTask(Tasks *[]ent.Task) error {
	if len(*Tasks) == 0 {
		return errors.New("Список задач пуст")
	}

	PrintTasks(Tasks)

	fmt.Println("Введите ID задачи, которую нужно отметить как выполненную:")
	var id int
	fmt.Scanln(&id)

	var found bool
	for i, t := range *Tasks {
		if t.ID == id {
			found = true
			(*Tasks)[i].Done = true
			break
		}
	}

	if !found {
		return errors.New("Задача с заданным ID не найдена")
	}

	if err := SaveTasks(*Tasks); err != nil {
		return errors.New("Ошибка сохранения задачи")
	}

	fmt.Println("Задача успешно отмечена как выполненная.")

	return nil
}

func DeleteTask(Tasks *[]ent.Task) error {
	if len(*Tasks) == 0 {
		return errors.New("Список задач пуст")
	}

	PrintTasks(Tasks)
	fmt.Println("Введите ID задачи для удаления:")
	
	var id int
	fmt.Scanln(&id)

	var found bool
	for i, t := range *Tasks {
		if t.ID == id {
			found = true
			*Tasks = append((*Tasks)[:i], (*Tasks)[i+1:]...)
			break
		}
	}

	if !found {
		return errors.New("Задача с заданным ID не найдена")
	}

	if err := SaveTasks(*Tasks); err != nil {
		return errors.New("Ошибка сохранения задачи")
	}

	fmt.Println("Задача успешно удалена.")

	return nil
}

func ReadTasks() ([]ent.Task, error) {
	db, err := sql.Open("sqlite3", tasksDB)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, text, done FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Tasks []ent.Task

	for rows.Next() {
		var id int
		var text string
		var done bool

		err := rows.Scan(&id, &text, &done)
		if err != nil {
			return nil, err
		}

		Tasks = append(Tasks, ent.NewTask(id, text, done))
	}

	return Tasks, nil
}

func SaveTasks(Tasks []ent.Task) error {
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

	for _, t := range Tasks {
		_, err = statement.Exec(t.ID, t.Text, t.Done)
		if err != nil {
			return err
		}
	}

	return nil
}
