// A program Team Project
//
// Purdue University -- CS18000 -- Spring 2024 -- Team Project
//
// Author: Purdue CS
// Version: Mar 31, 2024

package main

import "fmt"

// CommentIncompleteException is a custom error type in Go.
type CommentIncompleteException struct {
	Message string
}

// Implement the Error() method for the CommentIncompleteException to satisfy the error interface.
func (e *CommentIncompleteException) Error() string {
	return e.Message
}

// Example function that throws CommentIncompleteException
func NewComment(content string) error {
	if content == "" {
		return &CommentIncompleteException{Message: "Content is required"}
	}
	return nil
}

// Main function to demonstrate usage
func main() {
	err := NewComment("")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
