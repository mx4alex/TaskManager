package cli_application

import (
	"fmt"
	"TaskManager/internal/usecase"
	"errors"
	"bufio"
	"strings"
	"os"
)

type TaskCLI struct {
	taskInteractor *usecase.TaskInteractor
}

func NewTaskCLI(taskInteractor *usecase.TaskInteractor) *TaskCLI {
	return &TaskCLI{
		taskInteractor: taskInteractor,
	}
}

func (cli *TaskCLI) Run() error {
	var cmd string

	for {
		cli.CommandPrint()
		fmt.Scanln(&cmd)
		
		switch cmd {
		case "Create":
			err := cli.CreateTask()
			if err != nil {
				return err
			}
		case "Read":
			err := cli.PrintTasks()
			if err != nil {
				return err
			}
		case "Update":
			err := cli.UpdateTask()
			if err != nil {
				return err
			}
		case "Mark":
			err := cli.MarkTask()
			if err != nil {
				return err
			}
		case "Delete":
			err := cli.DeleteTask()
			if err != nil {
				return err
			}
		case "Exit":
			return nil
		default:
			return errors.New("Недопустимое действие.")
		}
	}
}

func (cli *TaskCLI) CommandPrint() {
	fmt.Println()
	fmt.Println("Выберите действие:")
	fmt.Println("Create - Создать задачу")
	fmt.Println("Read - Показать список задач")
	fmt.Println("Update - Обновить задачу")
	fmt.Println("Mark - Отметить задачу выполненной")
	fmt.Println("Delete - Удалить задачу")
	fmt.Println("Exit - Выйти")
	fmt.Println()
}

func (cli *TaskCLI) ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}


func (cli *TaskCLI) CreateTask() error {
	fmt.Println("Введите задачу:")
	text := cli.ReadInput()

	err := cli.taskInteractor.AddTask(text)
	if err != nil {
		return err
	}

	fmt.Println("Задача успешно создана:", text)
	return nil
}

func (cli *TaskCLI) PrintTasks() error {
	tasks, err := cli.taskInteractor.GetTasks()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("Список задач пуст")
		return nil
	}

	fmt.Println("Список задач:")
	for _, t := range tasks {
		if t.Done {
			fmt.Printf("%d. [x] %s\n", t.ID, t.Text)
		} else {
			fmt.Printf("%d. [ ] %s\n", t.ID, t.Text)
		}
	}

	return nil
}

func (cli *TaskCLI) UpdateTask() error {
	cli.PrintTasks()

	fmt.Println("Введите ID задачи для обновления:")
	var id int
	fmt.Scanf("%d\n", &id)

	fmt.Println("Введите новый текст задачи:")
	newText := cli.ReadInput()
	err := cli.taskInteractor.UpdateTask(id, newText)
	if err != nil {
		return err
	}

	fmt.Println("Задача успешно обновлена.")
	return nil
}

func (cli *TaskCLI) MarkTask() error {
	cli.PrintTasks()

	fmt.Println("Введите ID задачи, которую выполнили:")
	var id int
	fmt.Scanln(&id)

	err := cli.taskInteractor.MarkTask(id)
	if err != nil {
		return err
	}

	fmt.Println("Задача успешно отмечена как выполненная.")
	return nil
}

func (cli *TaskCLI) DeleteTask() error {
	fmt.Println("Введите ID задачи для удаления:")
	var id int
	fmt.Scanln(&id)

	err := cli.taskInteractor.DeleteTask(id)
	if err != nil {
		return err
	}

	fmt.Println("Задача успешно удалена.")
	return nil
}