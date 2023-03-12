package main

import (
	"fmt"
)

func updateTask(oldTask *Task, newTask *Task) {
	// Modify this function to create your Taskwarrior hook
}

func main() {
	oldTask, newTask, err := parseTasksFromStdin()

	if err != nil {
		fmt.Println("Error!")
	}

	updateTask(&oldTask, &newTask)

	newTaskStr, _ := newTask.JSON()
	fmt.Println(newTaskStr)
}
