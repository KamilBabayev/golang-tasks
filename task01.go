package main

import (
	"fmt"
	"os"
)

func main() {

	tasks := map[int]map[string]string{}

	task1 := map[string]string{
		"infra": "WRite ansible code to provision infra",
	}
	task2 := map[string]string{
		"aws_ecs": "Write terraform code to provision staging env",
	}

	tasks[0] = task1
	tasks[1] = task2
	fmt.Println(tasks)
	fmt.Println(tasks[1]["name"])

	for a, b := range tasks {
		fmt.Println(a, b)
	}

	mainMenu(tasks)

}

func mainMenu(tasks map[int]map[string]string) {
	for {
		fmt.Println("Task mgmt:")
		fmt.Println("1. List tasks")
		fmt.Println("2. Add new task")
		fmt.Println("3. Edit task")
		fmt.Println("4. Delete task")
		fmt.Println("5. Exit")
		var input string
		fmt.Scanln(&input)

		switch {
		case input == "1":
			fmt.Println("\nListing all tasks")
			for a, b := range tasks {
				fmt.Println(a, b)
			}
			fmt.Println("\n")
		case input == "2":
			fmt.Println("Add new task \n")
			var taskName, taskDesc string
			fmt.Scan(&taskName, &taskDesc)

			tempTask := map[string]string{
				taskName: taskDesc,
			}
			tasks[len(tasks)] = tempTask

		case input == "3":
			fmt.Println("Edit task \n")
			fmt.Println("Which task do you want to edit")
			for a, b := range tasks {
				fmt.Println(a, b)
			}
			fmt.Println("\n")

			var editTaskNum int
			fmt.Scan(&editTaskNum)
			fmt.Println("You are editing this task")
			fmt.Println(tasks[editTaskNum])
			fmt.Println("please enter updated version")
			var taskName, taskDesc string
			fmt.Scan(&taskName, &taskDesc)
			editTask := map[string]string{
				taskName: taskDesc,
			}
			tasks[editTaskNum] = editTask

			fmt.Println("Task updated as below")
			fmt.Println(tasks[editTaskNum])
			fmt.Println("\n")

		case input == "4":
			fmt.Println("delete task \n")
			fmt.Println("Which task do you want to delete ?")
			for a, b := range tasks {
				fmt.Println(a, b)
			}

			var deleteTask int
			fmt.Scan(&deleteTask)
			delete(tasks, deleteTask)
			fmt.Println("\n")

		case input == "5":
			fmt.Println("Exit from main menu")
			os.Exit(2)
		default:
			fmt.Println("please select operation to perform\n")

		}

	}

}
