package main

import (
	"os"
	"time"
)

type TaskStatus int

const (
	TaskTodo = iota
	TaskProgress
	TaskDone
)

func (t TaskStatus) String() string {
	switch t {
	case TaskTodo:
		return "Todo"
	case TaskProgress:
		return "In Progress"
	case TaskDone:
		return "Done"
	default:
		return "Unknown"
	}
}

type Task struct {
	ID          string     `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"UpdatedAt"`
}

var filename = "todolist.json"

func main() {
	checkFile(filename)
}

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return nil
		}
	}
	return nil
}
