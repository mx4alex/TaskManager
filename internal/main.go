package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const tasksFile = "tasks.csv"

type Task struct {
	ID   int
	Text string
	Done bool
}

func NewTask (id int, text string, done bool) Task {
	return Task {
		ID:   id,
		Text: text,
		Done: done,
	}
}

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

func CreateTask(Tasks *[]Task) {
	fmt.Println("Введите задачу:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	for _, t := range *Tasks {
		if t.Text == text {
			fmt.Println("Задача уже существует:", text)
			return
		}
	}

	*Tasks = append(*Tasks, NewTask(len(*Tasks) + 1, text, false))

	if err := SaveTasks(*Tasks); err != nil {
		fmt.Println("Ошибка сохранения задачи:", err)
		return
	}

	fmt.Println("Задача успешно создана:", text)
}

func PrintTasks(Tasks *[]Task) {
	if len(*Tasks) == 0 {
		fmt.Println("Список задач пуст.")
		return
	}

	fmt.Println("Список задач:")
	for _, t := range *Tasks {
		if t.Done {
			fmt.Printf("%d. [x] %s\n", t.ID, t.Text)
		} else {
			fmt.Printf("%d. [ ] %s\n", t.ID, t.Text)
		}
	}
}

func UpdateTask(Tasks *[]Task) {
	if len(*Tasks) == 0 {
		fmt.Println("Список задач пуст.")
		return
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
		fmt.Println("Задача с ID", id, "не найдена.")
		return
	}

	if err := SaveTasks(*Tasks); err != nil {
		fmt.Println("Ошибка сохранения задачи:", err)
		return
	}

	fmt.Println("Задача успешно обновлена.")
}

func MarkTask(Tasks *[]Task) {
	if len(*Tasks) == 0 {
		fmt.Println("Список задач пуст.")
		return
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
		fmt.Println("Задача с ID", id, "не найдена.")
		return
	}
	
	if err := SaveTasks(*Tasks); err != nil {
		fmt.Println("Ошибка сохранения задач:", err)
		return
	}

	fmt.Println("Задача успешно отмечена как выполненная.")
}

func DeleteTask(Tasks *[]Task) {
	if len(*Tasks) == 0 {
		fmt.Println("Список задач пуст.")
		return
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
		fmt.Println("Задача с ID", id, "не найдена.")
		return
	}

	if err := SaveTasks(*Tasks); err != nil {
		fmt.Println("Ошибка сохранения задач:", err)
		return
	}

	fmt.Println("Задача успешно удалена.")
}

func ReadTasks() ([]Task, error) {
	file, err := os.Open(tasksFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var Tasks []Task

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		
		id, _ := strconv.Atoi(record[0]) 
		done, _ := strconv.ParseBool(record[2])

		Tasks = append(Tasks, NewTask(id, record[1], done))
	}

	return Tasks, nil
}

func SaveTasks(Tasks []Task) error {
	file, err := os.OpenFile(tasksFile, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	
	writer := csv.NewWriter(file)

	for _, t := range Tasks {
		record := []string{strconv.Itoa(t.ID), t.Text, strconv.FormatBool(t.Done)}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	writer.Flush()

	return nil
}

func main() {
	var cmd string

	Tasks, err := ReadTasks()
	if err != nil {
		fmt.Println("Ошибка чтения данных о задачах:", err)
		return
	}

	for {
		CommandPrint()
		fmt.Scanln(&cmd)

		switch cmd {
		case "Create":
			CreateTask(&Tasks)
		case "Read":
			PrintTasks(&Tasks)
		case "Update":
			UpdateTask(&Tasks)
		case "Mark":
			MarkTask(&Tasks)
		case "Delete":
			DeleteTask(&Tasks)
		case "Exit":
			fmt.Println("Программа завершена.")
			return
		default:
			fmt.Println("Недопустимое действие.")
		}

		fmt.Println()
	}
}
