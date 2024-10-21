// A program Team Project
//
// Purdue University -- CS18000 -- Spring 2024 -- Team Project
//
// Author: Purdue CS
// Version: Mar 31, 2024

package main

import (
	"fmt"
	"time"
)

// LongRunningTask simulates a long-running background task
type LongRunningTask struct {
	done chan bool
}

// NewLongRunningTask creates a new instance of LongRunningTask
func NewLongRunningTask() *LongRunningTask {
	return &LongRunningTask{
		done: make(chan bool),
	}
}

// DoInBackground performs the long-running task in a goroutine
func (task *LongRunningTask) DoInBackground() {
	go func() {
		// Simulate a long-running task, e.g., network request or complex computation
		fmt.Println("Starting long-running task...")
		time.Sleep(5 * time.Second) // Simulate a 5-second delay (like Thread.sleep(5000) in Java)
		task.done <- true
	}()
}

// Done checks if the task is completed and processes the result
func (task *LongRunningTask) Done() {
	select {
	case <-task.done:
		// Task completed successfully, update UI or print a message
		fmt.Println("Task completed!")
	default:
		// Task is still running
		fmt.Println("Task is still running...")
	}
}

func main() {
	// Create a new long-running task
	task := NewLongRunningTask()

	// Start the long-running task in the background
	task.DoInBackground()

	// Continuously check if the task is done
	for {
		time.Sleep(1 * time.Second) // Simulate GUI event loop checking periodically
		task.Done()
		if <-task.done {
			break
		}
	}

	// Here you would normally update the GUI (if using a GUI framework)
	fmt.Println("Main function finished.")
}
