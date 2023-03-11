package main

import (
	"fmt"
	"os"
)

func updateTasks(oldTasks []Task, newTasks []Task) {
	// Get a environment variable with the log file
	logFile := os.Getenv("TASKWARRIOR_LOG")

	// Open the file
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		panic("Failed to open the log file")
	}

	defer file.Close()

	// Go through all changed tasks, and log the changes
	for index, _ := range newTasks {
		newTasksStr, _ := newTasks[index].JSON()
		oldTasksStr, _ := oldTasks[index].JSON()
		file.WriteString(fmt.Sprintf("The task changed from '%s' to '%s'.\n", oldTasksStr, newTasksStr))
	}
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
