// A program Team Project
//
// Purdue University -- CS18000 -- Spring 2024 -- Team Project
//
// Author: Purdue CS
// Version: Mar 31, 2024

package main

// PostRequirements defines the required methods for a Post in Go
type PostRequirements interface {
	Upvote()
	Downvote()
	Hide()
	Show()
}

// Post struct that will implement the PostRequirements interface
type Post struct {
	Title     string
	Content   string
	Upvotes   int
	Downvotes int
	Hidden    bool
}

// Upvote increases the upvotes count
func (p *Post) Upvote() {
	p.Upvotes++
}

// Downvote increases the downvotes count
func (p *Post) Downvote() {
	p.Downvotes++
}

// Hide hides the post
func (p *Post) Hide() {
	p.Hidden = true
}

// Show makes the post visible
func (p *Post) Show() {
	p.Hidden = false
}

// Example usage of Post implementing PostRequirements
func main() {
	post := &Post{Title: "Go Language", Content: "Go is great for concurrency", Upvotes: 0, Downvotes: 0, Hidden: false}

	// Applying actions on post
	post.Upvote()
	post.Hide()

	// Display the results
	if post.Hidden {
		println("Post is hidden")
	} else {
		println("Post is visible")
	}

	println("Upvotes:", post.Upvotes)
	println("Downvotes:", post.Downvotes)
}
