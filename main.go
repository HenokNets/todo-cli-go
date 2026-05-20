package main 

import (
  "fmt"
  "errors"
)

type Task struct {
  ID int
  Title string
  Done bool
}

var tasks []Task
var nextID = 1

func main () {
  //test everything 

  fmt.Println ("Testing the app")

  //test addTask 

  addTask ("Learn Go")
  addTask ("Build project")

  //test listTasks

  listTasks()

  //test removeTask

  err := removeTask (1)
  if err != nil {
    fmt.Println ("Error:", err)
  }

  //test listTasks again to see removal

  listTasks()

  //test error case 

  err = removeTask (99) 
  if err != nil {
    fmt.Println ("Expected error:", err)
  }
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

func removeTask (index int) error {
  idx := index - 1

  if idx < 0 || idx >= len (tasks) {
    return errors.New ("Index out of range")

  }

  removed := tasks [idx]
  tasks = append (tasks[:idx], tasks[idx+1:]...)

  fmt.Printf ("Removed task: %s\n", removed.Title)
  return nil //nil is zero type for interfaces (error is an interface type)



}