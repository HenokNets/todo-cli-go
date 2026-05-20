package main 

import "fmt"

type Task struct {
  ID int
  Title string
  Done bool
}

var tasks []Task
var nextID = 1

func main () {

}

func addTask (title string) Task {
  task := Task {
    ID: nextID,
    Title: title,
    Done: false,
  }

  tasks = append (tasks, task)
  nextID++
  fmt.Printf ("Added task: %s (ID: %d)\n", title, task.ID)
  return task
}

func listTasks () {

  if len (tasks) == 0 {
    fmt.Println ("No tasks available")
    return
  }
  fmt.Println ("\nYour tasks:")
  for i, task := range tasks {
    status := ""

    if task.Done {
      status = "✓"
    }

    fmt.Printf ("%d [%s] %s (ID: %d)", i+1, status, task.Title, task.ID)
  }
}