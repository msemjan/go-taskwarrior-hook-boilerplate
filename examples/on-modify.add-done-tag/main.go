package main

import (
	"fmt"
)

func updateTasks(oldTasks []Task, newTasks []Task) {
	for index, _ := range newTasks {
		if oldTasks[index].Status == "pending" && newTasks[index].Status == "done" {
			newTasks[index].Tags = append(newTasks[index].Tags, "Done")
		}
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
