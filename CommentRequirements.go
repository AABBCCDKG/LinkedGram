// A program Team Project
//
// Purdue University -- CS18000 -- Spring 2024 -- Team Project
//
// Author: Purdue CS
// Version: Mar 31, 2024

package main

// CommentRequirements is a Go interface that defines the required methods
type CommentRequirements interface {
	Upvote()
	Downvote()
	Hide()
	Show()
}

// Example of a struct that implements CommentRequirements
type Comment struct {
	upvotes   int
	downvotes int
	hidden    bool
}

// Implementing the Upvote method
func (c *Comment) Upvote() {
	c.upvotes++
}

// Implementing the Downvote method
func (c *Comment) Downvote() {
	c.downvotes++
}

// Implementing the Hide method
func (c *Comment) Hide() {
	c.hidden = true
}

// Implementing the Show method
func (c *Comment) Show() {
	c.hidden = false
}

// Main function to demonstrate usage
func main() {
	comment := &Comment{upvotes: 0, downvotes: 0, hidden: false}

	comment.Upvote()
	comment.Hide()
	comment.Show()

	// Example output
}
