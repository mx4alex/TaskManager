package main

import (
	"fmt"
	fnc "TaskManager/internal/taskfunc"
)

func main() {
	var cmd string

	Tasks, err := fnc.ReadTasks()
	if err != nil {
		fmt.Println("Ошибка чтения данных о задачах:", err)
		return
	}

	for {
		fnc.CommandPrint()
		fmt.Scanln(&cmd)

		switch cmd {
		case "Create":
			fnc.CreateTask(&Tasks)
		case "Read":
			fnc.PrintTasks(&Tasks)
		case "Update":
			fnc.UpdateTask(&Tasks)
		case "Mark":
			fnc.MarkTask(&Tasks)
		case "Delete":
			fnc.DeleteTask(&Tasks)
		case "Exit":
			fmt.Println("Программа завершена.")
			return
		default:
			fmt.Println("Недопустимое действие.")
		}

		fmt.Println()
	}
}
