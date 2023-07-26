package main

import (
	"fmt"
	"TaskManager/internal/app"
)

func main() {
	var cmd string
	
	taskManager := app.NewTaskManager()
	
	err := taskManager.ReadTasks()
	if err != nil {
		fmt.Println("Ошибка чтения данных о задачах:", err)
		// return
	}

	for {
		taskManager.CommandPrint()
		fmt.Scanln(&cmd)

		switch cmd {
		case "Create":
			err := taskManager.CreateTask()
			if err != nil {
				fmt.Println(err)
			}
		case "Read":
			err := taskManager.PrintTasks()
			if err != nil {
				fmt.Println(err)
			}
		case "Update":
			err := taskManager.UpdateTask()
			if err != nil {
				fmt.Println(err)
			}
		case "Mark":
			err := taskManager.MarkTask()
			if err != nil {
				fmt.Println(err)
			}
		case "Delete":
			err := taskManager.DeleteTask()
			if err != nil {
				fmt.Println(err)
			}
		case "Exit":
			err := taskManager.SaveTasks()
			if err != nil {
				fmt.Println("Ошибка сохранения задач:", err)
			}
			fmt.Println("Программа завершена.")
			return
		default:
			fmt.Println("Недопустимое действие.")
		}

		fmt.Println()
	}
}