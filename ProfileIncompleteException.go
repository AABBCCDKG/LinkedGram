// A program Team Project
//
// Purdue University -- CS18000 -- Spring 2024 -- Team Project
//
// Author: Purdue CS
// Version: Mar 31, 2024

package main

import "fmt"

// ProfileIncompleteException represents a custom error in Go
type ProfileIncompleteException struct {
	Message string
}

// Implementing the Error() method for ProfileIncompleteException to satisfy the error interface
func (e *ProfileIncompleteException) Error() string {
	return e.Message
}

// Example function that throws ProfileIncompleteException
func createProfile(name string) error {
	if name == "" {
		return &ProfileIncompleteException{Message: "Profile name is required"}
	}
	// Proceed with profile creation
	return nil
}

// Main function to demonstrate usage
func main() {
	err := createProfile("")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
