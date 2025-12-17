package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
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
	ID          uuid.UUID  `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

var filename = "todolist.json"

func main() {
	if err := checkFile(filename); err != nil {
		log.Fatal(err)
	}

	arguments := os.Args

	if len(arguments) == 1 {
		return
	}

	comands := arguments[1]

	tasks := []Task{}
	if err := loadTasks(&tasks); err != nil {
		log.Fatal(err)
	}
	switch comands {
	case "add":
		handleAdd(arguments, tasks)
	case "list":
		handleList(arguments, tasks)
	}
}

func handleAdd(arguments []string, tasks []Task) {
	if len(arguments) < 3 {
		fmt.Println("add <title> [description]")
		return
	}

	var description string

	if len(arguments) > 3 {
		description = arguments[3]
	} else {
		description = "(none)"
	}

	taskId, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}

	task := &Task{
		ID:          taskId,
		Description: description,
		Status:      TaskTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, *task)
	if err := saveTasks(tasks); err != nil {
		log.Fatal(err)
	}

	fmt.Println("List added")
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

func handleList(arguments []string, tasks []Task) {
	timeFormat := "jan _2 15:04:05 2006"
	filterByStatus := len(arguments) > 2
	var status TaskStatus

	if filterByStatus {
		switch arguments[2] {
		case "todo":
			status = TaskTodo
		case "in-progress":
			status = TaskProgress
		case "done":
			status = TaskDone
		default:
			filterByStatus = false
		}
	}

	for i, task := range tasks {
		if !filterByStatus || status == task.Status {
			fmt.Printf("List. %v\nDescription: %v\nStatus: %v\nCreated At: %v\nUpdate At: %v\n", i+1,
				task.Description,
				task.Status,
				task.CreatedAt.Format(timeFormat),
				task.UpdatedAt.Format(timeFormat))

			fmt.Println()
		}
	}
}

func loadTasks(tasks *[]Task) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	if len(file) > 0 {
		if err := json.Unmarshal(file, &tasks); err != nil {
			return err
		}
	}

	return nil
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		return err
	}

	os.WriteFile(filename, data, 0o644)
	return nil
}
