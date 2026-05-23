package main

import (
	"errors"
	"fmt"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func (t Task) String() string {
	status := " "
	if t.Done {
		status = "✓"
	}

	return fmt.Sprintf("[%s] %d. %s", status, t.ID, t.Title)
}

func addTask(title string) Task {
	task := Task{
		ID:    nextID,
		Title: title,
		Done:  false,
	}

	tasks = append(tasks, task)
	nextID++
	fmt.Printf("Added: %s\n", task)
	return task
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available")
		return
	}
	fmt.Println("\nID  | Status | Title")
	fmt.Println("----------------------")
	for _, task := range tasks {
		status := "✗"
		if task.Done {
			status = "✓"
		}
		fmt.Printf("%-3d | %-6s | %s\n", task.ID, status, task.Title)
	}
}

func removeTask(index int) error {
	idx := index - 1

	if idx < 0 || idx >= len(tasks) {
		return errors.New("index out of range")
	}

	removed := tasks[idx]
	tasks = append(tasks[:idx], tasks[idx+1:]...)

	fmt.Printf("Removed: %s\n", removed)
	return nil
}

func doneTask(index int) error {
	idx := index - 1

	if idx < 0 || idx >= len(tasks) {
		return errors.New("index out of range")
	}

	tasks[idx].Done = true
	fmt.Printf("Marked done: %s\n", tasks[idx])
	return nil
}