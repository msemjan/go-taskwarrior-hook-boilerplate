package main

import (
	"fmt"
)

func updateTask(oldTask *Task, newTask *Task) {
	if *(oldTask.Status) == "pending" && *(newTask.Status) == "completed" {
		if newTask.Tags != nil {
			newTask.Tags = append(newTask.Tags, "Done")
		} else {
			newTask.Tags = []string{"Done"}
		}
	}
}

func main() {
	oldTask, newTask, err := parseTaskFromStdin()

	if err != nil {
		fmt.Println("Error!")
	}

	updateTask(&oldTask, &newTask)

	newTaskStr, _ := newTask.JSON()
	fmt.Println(newTaskStr)
}
