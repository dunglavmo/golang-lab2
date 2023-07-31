package main

import (
	"fmt"
)

func main() {
	cache := NewTaskCache()

	// Add tasks
	task1 := Task{ID: 1, Name: "Task 1", Status: 0, CreatedAt: "2023-07-20"}
	task2 := Task{ID: 2, Name: "Task 2", Status: 1, CreatedAt: "2023-07-21"}
	cache.AddTask(task1)
	cache.AddTask(task2)

	// List all tasks
	tasks, err := cache.ListTasks()
	if err != nil {
		fmt.Println("Error listing tasks:", err)
	} else {
		fmt.Println("All Tasks:")
		for _, t := range tasks {
			fmt.Printf("ID: %d, Name: %s, Status: %d, CreatedAt: %s\n", t.ID, t.Name, t.Status, t.CreatedAt)
		}
	}

	// Update a task
	task1.Name = "Updated Task 1"
	cache.UpdateTask(task1)

	// Get a specific task
	taskID := 2
	task, err := cache.GetTask(taskID)
	if err != nil {
		fmt.Println("Error getting task:", err)
	} else {
		fmt.Printf("Task with ID %d: Name: %s, Status: %d, CreatedAt: %s\n", task.ID, task.Name, task.Status, task.CreatedAt)
	}

	// Delete a task
	cache.DeleteTask(taskID)

	// List all tasks after delete
	tasks, err = cache.ListTasks()
	if err != nil {
		fmt.Println("Error listing tasks:", err)
	} else {
		fmt.Println("All Tasks after delete:")
		for _, t := range tasks {
			fmt.Printf("ID: %d, Name: %s, Status: %d, CreatedAt: %s\n", t.ID, t.Name, t.Status, t.CreatedAt)
		}
	}
}
