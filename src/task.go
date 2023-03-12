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
	Depends     *string  `json:"depends,omitempty"`
	Start       *string  `json:"start,omitempty"`
	Until       *string  `json:"until,omitempty"`
	Wait        *string  `json:"wait,omitempty"`
	End         *string  `json:"end,omitempty"`
	Entry       *string  `json:"due,omitempty"`
	Modified    *string  `json:"modified,omitempty"`
	Scheduled   *string  `json:"scheduled,omitempty"`
	Recur       *string  `json:"recur,omitempty"`
	Mask        *string  `json:"mask,omitempty"`
	Imask       *string  `json:"imask,omitempty"`
	Priority    *string  `json:"priority,omitempty"`
	Parent      *string  `json:"parent,omitempty"`
	Project     *string  `json:"project,omitempty"`
	Status      *string  `json:"status,omitempty"`
	Uuid        string   `json:"uuid"`
	Tags        []string `json:"tags,omitempty"`
	Annotation  []string `json:"annotation,omitempty"`
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

func parseOneTaskFromStdin() (Task, error) {
	reader := bufio.NewReader(os.Stdin)

	newTask := Task{}

	line, err := reader.ReadString('\n')

	if err != nil {
		return newTask, err
	}

	err = json.Unmarshal([]byte(line), &newTask)

	return newTask, err
}

func parseTasksFromStdin() (Task, Task, error) {
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
