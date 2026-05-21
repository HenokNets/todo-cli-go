package main 

import (
  "fmt"
  "errors"
  "encoding/json"
  "os"
)

type Task struct {
  ID int `json:"id"`
  Title string `json:"title"`
  Done bool	`json:"done"`
}

func (t Task) String() string {
	status := " "
	if t.Done {
		status = "✓"
	}

	return fmt.Sprintf ("[%s] %d. %s", status, t.ID, t.Title)
}

var tasks []Task
var nextID = 1

func main() {
    fmt.Println("Todo App")

    //load tasks from file on startup
    err := loadTasks("tasks.json")
    if err != nil {
        fmt.Println("Error loading tasks:", err)
    }

    //show existing tasks
    listTasks()

    //add new tasks
    addTask("Learn Go")
    addTask("Build project")

    //show updated list
    listTasks()

    //remove a task
    err = removeTask(1)
    if err != nil {
        fmt.Println("Error:", err)
    }

    //show after removal
    listTasks()

    //save to file before exiting
    err = saveTasks("tasks.json")
    if err != nil {
        fmt.Println("Error saving tasks:", err)
    }

    fmt.Println("Tasks saved.")
}

func addTask (title string) Task {
  task := Task {
    ID: nextID,
    Title: title,
    Done: false,
  }

  tasks = append (tasks, task)
  nextID++
  fmt.Printf ("Added: %s\n", task)
  return task
}

func listTasks () {

  if len (tasks) == 0 {
    fmt.Println ("No tasks available")
    return
  }
  fmt.Println ("\nYour tasks:")
  for _, task := range tasks {
    fmt.Println(task)
  }
}

func removeTask (index int) error {
  idx := index - 1

  if idx < 0 || idx >= len (tasks) {
    return errors.New ("Index out of range")

  }

  removed := tasks [idx]
  tasks = append (tasks[:idx], tasks[idx+1:]...)

  fmt.Printf ("Removed: %s\n", removed)
  return nil //nil is zero type for interfaces (error is an interface type)



}

func saveTasks (filename string) error {
	data, err := json.MarshalIndent (tasks, "", "  ")
  	if err != nil {
		return err
  	}

	err = os.WriteFile (filename, data, 0644)
	if err != nil {
		return err
	}
	return nil


}

func loadTasks (filename string) error {
	data, err := os.ReadFile (filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf ("failed to read file: %w", err)
	}

	err = json.Unmarshal (data, &tasks)
	if err != nil {
		return fmt.Errorf ("failed to unmarshal tasks: %w", err)
	}

	for _, task := range tasks {
		if task.ID >= nextID {
			nextID = task.ID + 1
		}
	}

	return nil
}
