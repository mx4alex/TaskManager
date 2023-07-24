package taskfunc

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"errors"
	ent "TaskManager/internal/entity"
)

const tasksFile = "tasks.csv"

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
	file, err := os.OpenFile(tasksFile, os.O_CREATE|os.O_RDWR, 0777) 
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var Tasks []ent.Task

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

		Tasks = append(Tasks, ent.NewTask(id, record[1], done))
	}

	return Tasks, nil
}

func SaveTasks(Tasks []ent.Task) error {
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
