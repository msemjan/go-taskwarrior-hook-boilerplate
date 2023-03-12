package main

import (
	"fmt"
)

func updateTask(newTask *Task) {
	// Modify this function to create your Taskwarrior hook
}

func main() {
	newTask, err := parseOneTaskFromStdin()

	if err != nil {
		fmt.Println("Error!")
	}

	updateTask(&newTask)

	newTaskStr, _ := newTask.JSON()
	fmt.Println(newTaskStr)
}
