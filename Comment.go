// A program Team Project
//
// Purdue University -- CS18000 -- Spring 2024 -- Team Project
//
// Author: Purdue CS
// Version: Mar 31, 2024

package main

import (
	"errors"
	"fmt"
)

// Post struct simulates the Post class in Java
type Post struct {
	Title     string
	Content   string
	Author    string
	Hidden    bool
	ImageURL  string
	Upvotes   int
	Downvotes int
}

// CommentRequirements interface defines the required methods for a comment
type CommentRequirements interface {
	Upvote()
	Downvote()
	Hide()
	Show()
}

// Comment struct implements CommentRequirements and extends Post
type Comment struct {
	Post      Post
	Content   string
	Author    string
	Upvotes   int
	Downvotes int
	Hidden    bool
}

// CommentIncompleteException simulates the custom exception in Java
type CommentIncompleteException struct {
	message string
}

func (e *CommentIncompleteException) Error() string {
	return e.message
}

// NewComment creates a new Comment and checks for completeness
func NewComment(post Post, content, author string, hidden bool, upvotes, downvotes int) (*Comment, error) {
	if content == "" {
		return nil, &CommentIncompleteException{message: "Content is required"}
	}
	if author == "" {
		return nil, &CommentIncompleteException{message: "Author is required"}
	}
	return &Comment{
		Post:      post,
		Content:   content,
		Author:    author,
		Hidden:    hidden,
		Upvotes:   upvotes,
		Downvotes: downvotes,
	}, nil
}

// Upvote increases the upvotes of the comment
func (c *Comment) Upvote() {
	c.Upvotes++
}

// Downvote increases the downvotes of the comment
func (c *Comment) Downvote() {
	c.Downvotes++
}

// Hide hides the comment
func (c *Comment) Hide() {
	c.Hidden = true
}

// Show makes the comment visible
func (c *Comment) Show() {
	c.Hidden = false
}

// String returns the string representation of the Comment
func (c *Comment) String() string {
	return fmt.Sprintf("\nAuthor: %s\nContent: %s\nUpvotes: %d\nDownvotes: %d", c.Author, c.Content, c.Upvotes, c.Downvotes)
}

// Getters for Comment fields
func (c *Comment) GetContent() string {
	return c.Content
}

func (c *Comment) GetAuthor() string {
	return c.Author
}

func (c *Comment) IsHidden() bool {
	return c.Hidden
}

func (c *Comment) GetUpvotes() int {
	return c.Upvotes
}

func (c *Comment) GetDownvotes() int {
	return c.Downvotes
}

func (c *Comment) GetPost() Post {
	return c.Post
}

// Setters for Comment fields
func (c *Comment) SetContent(content string) {
	c.Content = content
}

func (c *Comment) SetAuthor(author string) {
	c.Author = author
}

func (c *Comment) SetHidden(hidden bool) {
	c.Hidden = hidden
}

func (c *Comment) SetUpvotes(upvotes int) {
	c.Upvotes = upvotes
}

func (c *Comment) SetDownvotes(downvotes int) {
	c.Downvotes = downvotes
}

func (c *Comment) SetPost(post Post) {
	c.Post = post
}

// Example usage of the Comment struct
func main() {
	post := Post{
		Title:   "Go Programming",
		Content: "Go is a statically typed, compiled programming language.",
		Author:  "Jane Doe",
		Hidden:  false,
	}

	comment, err := NewComment(post, "Great post!", "John Doe", false, 0, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(comment.String())
	comment.Upvote()
	fmt.Println("After upvote:", comment.String())
}
