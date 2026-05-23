package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"strconv"
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

var tasks []Task
var nextID = 1

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	err := loadTasks("tasks.json")
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: no title provided")
			return
		}
		title := strings.Join(os.Args[2:], " ")
		addTask(title)
		err = saveTasks("tasks.json")
		if err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}
		fmt.Printf("Saved %d tasks\n", len(tasks))
	case "list":
		listTasks()
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Error: no ID provided")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: invalid ID")
			return
		}
		err = removeTask(id)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		err = saveTasks("tasks.json")
		if err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}
		fmt.Printf("Saved %d tasks\n", len(tasks))
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Error: no ID provided")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: invalid ID")
			return
		}
		err = doneTask(id)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		err = saveTasks("tasks.json")
		if err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}
		fmt.Printf("Saved %d tasks\n", len(tasks))
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  todo add <title>")
	fmt.Println("  todo list")
	fmt.Println("  todo remove <id>")
	fmt.Println("  todo done <id>")
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
	fmt.Println("\nYour tasks:")
	for _, task := range tasks {
		fmt.Println(task)
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

func saveTasks(filename string) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("couldn't marshal tasks: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("couldn't write data: %w", err)
	}
	return nil
}

func loadTasks(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("failed to read file: %w", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return fmt.Errorf("failed to decode tasks: %w", err)
	}

	for _, task := range tasks {
		if task.ID >= nextID {
			nextID = task.ID + 1
		}
	}

	return nil
}