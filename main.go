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

	runCommand(os.Args[1], os.Args[2:])
}

func runCommand(cmd string, args []string) {
	switch cmd {
	case "add":
		if len(args) < 1 {
			fmt.Println("Error: no title provided")
			return
		}
		title := strings.Join(args, " ")
		addTask(title)
		saveIfModified()

	case "list":
		listTasks()

	case "remove":
		id, ok := parseID(args)
		if !ok {
			return
		}
		err := removeTask(id)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		saveIfModified()

	case "done":
		id, ok := parseID(args)
		if !ok {
			return
		}
		err := doneTask(id)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		saveIfModified()

	case "clear":
		count := clearDoneTasks()
		if count == 0 {
			fmt.Println("No completed tasks to clear")
		} else {
			fmt.Printf("Removed %d completed task(s)\n", count)
		}
		saveIfModified()

	case "filter":
		if len(args) < 1 {
			fmt.Println("Error: specify 'done' or 'undone'")
			return
		}
		filterTasks(args[0])

	default:
		printUsage()
	}
}

func parseID(args []string) (int, bool) {
	if len(args) < 1 {
		fmt.Println("Error: no ID provided")
		return 0, false
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Error: invalid ID")
		return 0, false
	}
	return id, true
}

func saveIfModified() {
	err := saveTasks("tasks.json")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	fmt.Printf("Saved %d tasks\n", len(tasks))
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