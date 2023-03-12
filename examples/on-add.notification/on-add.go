package main

import (
	"fmt"
	"log"
	"os/exec"
)

func updateTask(newTask *Task) {
	notification := fmt.Sprintf("Task '%s' was added", newTask.Description)
	cmd := exec.Command("notify-send", "-a", "Taskwarrior", "-t", "360000", notification)

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
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
