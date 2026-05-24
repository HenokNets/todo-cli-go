package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	case "clear":
		count := clearDoneTasks()
		if count == 0 {
			fmt.Println("No completed tasks to clear")
		} else {
			fmt.Printf("Removed %d completed task(s)\n", count)
		}
		err = saveTasks("tasks.json")
		if err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}
		fmt.Printf("Saved %d tasks\n", len(tasks))
	case "filter":
		if len(os.Args) < 3 {
			fmt.Println("Error: specify 'done' or 'undone'")
			return
		}
		filterTasks(os.Args[2])
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
	fmt.Println("  todo clear")
	fmt.Println("  todo filter <done|undone>")
}