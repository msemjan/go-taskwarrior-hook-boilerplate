package main

import (
	"bufio"
	"encoding/json"
	"os"
)

type Task struct {
	Id          int      `json:id`
	Description string   `json:description`
	Due         string   `json:due`
	Entry       string   `json:due`
	Modified    string   `json:modified`
	Priority    string   `json:priority`
	Project     string   `json:project`
	Status      string   `json:status`
	Uuid        string   `json:uuid`
	Tags        []string `json:tags`
	Urgency     float64  `json:urgency`
}

func (t *Task) JSON() (string, error) {
	jsonStr, err := json.Marshal(t)

	return string(jsonStr), err
}

func TaskToJSON(t []Task) (string, error) {
	jsonStr, err := json.Marshal(t)

	return string(jsonStr), err
}

func parseTasksFromStdin() ([]Task, []Task, error) {
	reader := bufio.NewReader(os.Stdin)

	oldTasks := []Task{}
	newTasks := []Task{}

	line, err := reader.ReadString('\n')

	if err != nil {
		return oldTasks, newTasks, err
	}

	err = json.Unmarshal([]byte(line), &oldTasks)

	if err != nil {
		return oldTasks, newTasks, err
	}

	line, err = reader.ReadString('\n')

	if err != nil {
		return oldTasks, newTasks, err
	}

	err = json.Unmarshal([]byte(line), &newTasks)

	return oldTasks, newTasks, err
}
