package main

import (
	"os"
	"testing"
)

func resetState() {
	tasks = []Task{}
	nextID = 1
}

func TestAddTask(t *testing.T) {
	resetState()

	task := addTask("Learn Go")

	if task.ID != 1 {
		t.Errorf("expected ID 1, got %d", task.ID)
	}
	if task.Title != "Learn Go" {
		t.Errorf("expected title 'Learn Go', got '%s'", task.Title)
	}
	if task.Done != false {
		t.Errorf("expected Done to be false, got true")
	}
	if len(tasks) != 1 {
		t.Errorf("expected 1 task in slice, got %d", len(tasks))
	}
}

func TestAddMultipleTasks(t *testing.T) {
	resetState()

	addTask("First")
	addTask("Second")
	addTask("Third")

	if len(tasks) != 3 {
		t.Errorf("expected 3 tasks, got %d", len(tasks))
	}
	if tasks[0].ID != 1 {
		t.Errorf("expected first task ID 1, got %d", tasks[0].ID)
	}
	if tasks[2].ID != 3 {
		t.Errorf("expected third task ID 3, got %d", tasks[2].ID)
	}
	if nextID != 4 {
		t.Errorf("expected nextID 4, got %d", nextID)
	}
}

func TestRemoveTask(t *testing.T) {
	resetState()

	addTask("First")
	addTask("Second")
	addTask("Third")

	err := removeTask(2) // remove second task
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(tasks) != 2 {
		t.Errorf("expected 2 tasks after removal, got %d", len(tasks))
	}
	if tasks[0].Title != "First" {
		t.Errorf("expected 'First', got '%s'", tasks[0].Title)
	}
	if tasks[1].Title != "Third" {
		t.Errorf("expected 'Third', got '%s'", tasks[1].Title)
	}
}

func TestRemoveTaskOutOfRange(t *testing.T) {
	resetState()

	addTask("First")

	err := removeTask(0) 
	if err == nil {
		t.Error("expected error for index 0, got nil")
	}

	err = removeTask(2) // only one task exists
	if err == nil {
		t.Error("expected error for index 2, got nil")
	}
}

func TestDoneTask(t *testing.T) {
	resetState()

	addTask("First")
	err := doneTask(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if tasks[0].Done != true {
		t.Error("expected task to be marked done")
	}
}

func TestDoneTaskAlreadyDone(t *testing.T) {
	resetState()

	addTask("First")
	doneTask(1)
	err := doneTask(1)
	if err == nil {
		t.Error("expected error for already done task, got nil")
	}
}

func TestLoadTasks(t *testing.T) {
	resetState()

	//create a temporary tasks file
	testFile := "test_tasks.json"
	defer os.Remove(testFile)

	addTask("Persisted task")
	addTask("Another task")
	saveTasks(testFile)

	//reset and load
	resetState()
	err := loadTasks(testFile)
	if err != nil {
		t.Errorf("unexpected error loading tasks: %v", err)
	}
	if len(tasks) != 2 {
		t.Errorf("expected 2 loaded tasks, got %d", len(tasks))
	}
	if tasks[0].Title != "Persisted task" {
		t.Errorf("expected 'Persisted task', got '%s'", tasks[0].Title)
	}
	if tasks[1].Title != "Another task" {
		t.Errorf("expected 'Another task', got '%s'", tasks[1].Title)
	}
	if nextID != 3 {
		t.Errorf("expected nextID 3 after load, got %d", nextID)
	}
}

func TestLoadTasksFileNotExist(t *testing.T) {
	resetState()

	err := loadTasks("nonexistent_file.json")
	if err != nil {
		t.Errorf("expected no error for missing file, got %v", err)
	}
	if len(tasks) != 0 {
		t.Errorf("expected empty tasks for missing file, got %d", len(tasks))
	}
}

func TestLoadTasksCorruptedJSON(t *testing.T) {
	resetState()

	testFile := "corrupted_tasks.json"
	defer os.Remove(testFile)

	//write invalid json
	err := os.WriteFile(testFile, []byte("this is not json"), 0644)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	err = loadTasks(testFile)
	if err == nil {
		t.Error("expected error for corrupted JSON, got nil")
	}
}