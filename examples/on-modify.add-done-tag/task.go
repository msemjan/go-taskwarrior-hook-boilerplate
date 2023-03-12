package main

import (
	"bufio"
	"encoding/json"
	"os"
)

type Task struct {
	Id          int      `json:"id"`
	Description string   `json:"description"`
	Due         *string  `json:"due,omitempty"`
	End         *string  `json:"end,omitempty"`
	Entry       *string  `json:"due,omitempty"`
	Modified    *string  `json:"modified,omitempty"`
	Priority    *string  `json:"priority,omitempty"`
	Project     *string  `json:"project,omitempty"`
	Status      *string  `json:"status,omitempty"`
	Uuid        string   `json:"uuid"`
	Tags        []string `json:"tags,omitempty"`
	Urgency     *float64 `json:"urgency,omitempty"`
}

func (t *Task) JSON() (string, error) {
	jsonStr, err := json.Marshal(t)

	return string(jsonStr), err
}

func TasksToJSON(t []Task) (string, error) {
	jsonStr, err := json.Marshal(t)

	return string(jsonStr), err
}

func parseTaskFromStdin() (Task, Task, error) {
	reader := bufio.NewReader(os.Stdin)

	oldTask := Task{}
	newTask := Task{}

	line, err := reader.ReadString('\n')

	if err != nil {
		return oldTask, newTask, err
	}

	err = json.Unmarshal([]byte(line), &oldTask)

	if err != nil {
		return oldTask, newTask, err
	}

	line, err = reader.ReadString('\n')

	if err != nil {
		return oldTask, newTask, err
	}

	err = json.Unmarshal([]byte(line), &newTask)

	return oldTask, newTask, err
}
