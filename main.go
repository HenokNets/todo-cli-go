package main 

import (
  "fmt"
  "errors"
  "encoding/json"
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

  data, err := json.Marshal (tasks)
  if err != nil {
	fmt.Println (err)
  }
  
  fmt.Println (data)



  jsonData := []byte(`
	[{
  		"id": 122,
		"title": "study go",
		"done": true
	}]
`)
var loadedTasks []Task
  err = json.Unmarshal (jsonData, &loadedTasks)
  if err != nil {
	fmt.Println("Unmarshal error:", err)
  }
 fmt.Println ("Unmarshaled:", loadedTasks)
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

