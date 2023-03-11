package main

import (
	"fmt"
)

func updateTasks(oldTasks []Task, newTasks []Task) {
	// Modify this function to create your Taskwarrior hook
}

func main() {
	oldTasks, newTasks, err := parseTasksFromStdin()

	if err != nil {
		fmt.Println("Error!")
	}

	updateTasks(oldTasks, newTasks)

	newTasksStr, _ := TaskToJSON(newTasks)
	fmt.Println(newTasksStr)
}
