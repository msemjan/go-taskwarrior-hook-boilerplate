package main

import (
	"fmt"
	"os"
)

func updateTask(oldTask Task, newTask Task) {
	// Get a environment variable with the log file
	logFile := os.Getenv("TASKWARRIOR_LOG")

	// Open the file
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		panic("Failed to open the log file")
	}

	defer file.Close()

	// Go through all changed tasks, and log the changes
	newTaskStr, _ := newTask.JSON()
	oldTaskStr, _ := oldTask.JSON()
	file.WriteString(fmt.Sprintf("The task changed from '%s' to '%s'.\n", oldTaskStr, newTaskStr))
}

func main() {
	oldTask, newTask, err := parseTaskFromStdin()

	if err != nil {
		fmt.Println("Error!")
	}

	updateTask(oldTask, newTask)

	newTaskStr, _ := newTask.JSON()
	fmt.Println(newTaskStr)
}
