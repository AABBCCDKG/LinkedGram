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

// Comment struct to represent a comment
type Comment struct {
	Post      *Post
	Message   string
	Writer    string
	Hidden    bool
	Upvotes   int
	Downvotes int
}

// Post struct to represent a post
type Post struct {
	Title      string
	Content    string
	Author     string
	ImageURL   string
	Upvotes    int
	Downvotes  int
	Hidden     bool
	Comments   []Comment
	CommentsStr string
}

// PostIncompleteException simulates custom exception
var PostIncompleteException = errors.New("Post data incomplete")

// CommentIncompleteException simulates custom exception for comments
var CommentIncompleteException = errors.New("Comment data incomplete")

// NewPost creates a new Post with required parameters
func NewPost(title, content, author, imageURL string, upvotes, downvotes int, hidden bool, line string) (*Post, error) {
	if title == "" {
		return nil, errors.New("Title is required")
	}
	if content == "" {
		return nil, errors.New("Content is required")
	}
	if author == "" {
		return nil, errors.New("Author is required")
	}
	if imageURL == "" {
		imageURL = "No image"
	}

	return &Post{
		Title:       title,
		Content:     content,
		Author:      author,
		ImageURL:    imageURL,
		Upvotes:     upvotes,
		Downvotes:   downvotes,
		Hidden:      hidden,
		Comments:    []Comment{},
		CommentsStr: line,
	}, nil
}

// AddComment adds a comment to a post
func (p *Post) AddComment(writer, message string) error {
	if writer == "" {
		return CommentIncompleteException
	}
	if message == "" {
		return CommentIncompleteException
	}

	comment := Comment{
		Post:    p,
		Message: message,
		Writer:  writer,
		Hidden:  false,
		Upvotes: 0,
		Downvotes: 0,
	}
	p.Comments = append(p.Comments, comment)
	return nil
}

// ShowComments displays all comments of a post
func (p *Post) ShowComments() {
	for _, comment := range p.Comments {
		fmt.Println(comment)
	}
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

// ToString returns a string representation of the post
func (p *Post) ToString() string {
	return fmt.Sprintf("Title: %s\nContent: %s\nAuthor: %s\nImage URL: %s\nUpvotes: %d\nDownvotes: %d\nComments: %s\n", 
		p.Title, p.Content, p.Author, p.ImageURL, p.Upvotes, p.Downvotes, p.CommentsStr)
}

// Getters for post fields
func (p *Post) GetTitle() string {
	return p.Title
}

func (p *Post) GetContent() string {
	return p.Content
}

func (p *Post) GetAuthor() string {
	return p.Author
}

func (p *Post) GetUpvotes() int {
	return p.Upvotes
}

func (p *Post) GetDownvotes() int {
	return p.Downvotes
}

func (p *Post) IsHidden() bool {
	return p.Hidden
}

func (p *Post) GetImageURL() string {
	return p.ImageURL
}

// Setters for post fields
func (p *Post) SetTitle(title string) {
	p.Title = title
}

func (p *Post) SetContent(content string) {
	p.Content = content
}

func (p *Post) SetAuthor(author string) {
	p.Author = author
}

func (p *Post) SetImageURL(imageURL string) {
	p.ImageURL = imageURL
}

func (p *Post) SetUpvotes(upvotes int) {
	p.Upvotes = upvotes
}

func (p *Post) SetDownvotes(downvotes int) {
	p.Downvotes = downvotes
}

func (p *Post) SetHidden(hidden bool) {
	p.Hidden = hidden
}

// Main function to demonstrate usage
func main() {
	post, err := NewPost("Go Language", "Go is great for concurrency", "John", "golang.jpg", 10, 2, false, "Example comment")
	if err != nil {
		fmt.Println("Error creating post:", err)
		return
	}

	fmt.Println(post.ToString())

	// Add a comment to the post
	err = post.AddComment("Jane", "I agree!")
	if err != nil {
		fmt.Println("Error adding comment:", err)
		return
	}

	// Show comments
	post.ShowComments()

	// Upvote and downvote the post
	post.Upvote()
	post.Downvote()

	// Display updated post info
	fmt.Println(post.ToString())
}
