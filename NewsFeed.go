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
	"strings"
)

// Post struct simulates the Post class
type Post struct {
	Title   string
	Content string
	Author  string
}

// String method for Post struct to return a formatted string
func (p Post) String() string {
	return fmt.Sprintf("Title: %s\nContent: %s\nAuthor: %s\n", p.Title, p.Content, p.Author)
}

// NewsFeed struct, similar to the Java version
type NewsFeed struct {
	PostList []Post
	Owner    string
}

// NewNewsFeed creates a new NewsFeed instance
func NewNewsFeed() *NewsFeed {
	return &NewsFeed{
		PostList: []Post{},
	}
}

// AddPost adds a post to the NewsFeed
func (n *NewsFeed) AddPost(post Post) {
	n.PostList = append(n.PostList, post)
}

// RemovePost removes a post from the NewsFeed
func (n *NewsFeed) RemovePost(post Post) error {
	for i, p := range n.PostList {
		if p.Title == post.Title {
			n.PostList = append(n.PostList[:i], n.PostList[i+1:]...)
			return nil
		}
	}
	return errors.New("post not found")
}

// GetPost retrieves a post by its title
func (n *NewsFeed) GetPost(title string) *Post {
	for _, post := range n.PostList {
		if post.Title == title {
			return &post
		}
	}
	return nil
}

// ViewNewsFeed prints all posts in the NewsFeed
func (n *NewsFeed) ViewNewsFeed() error {
	if len(n.PostList) == 0 {
		return errors.New("no posts available")
	}
	for _, post := range n.PostList {
		fmt.Println(post)
	}
	return nil
}

// GetOwner returns the owner of the NewsFeed
func (n *NewsFeed) GetOwner() string {
	return n.Owner
}

// ToString returns the string representation of the NewsFeed
func (n *NewsFeed) ToString() string {
	var sb strings.Builder
	for _, post := range n.PostList {
		sb.WriteString(post.String())
		sb.WriteString("\n")
	}
	return sb.String()
}

// Main function demonstrating usage
func main() {
	// Create a new NewsFeed
	newsFeed := NewNewsFeed()

	// Create some posts
	post1 := Post{Title: "Go Programming", Content: "Go is an open-source language.", Author: "John"}
	post2 := Post{Title: "Concurrency in Go", Content: "Go provides great support for concurrency.", Author: "Alice"}

	// Add posts to the news feed
	newsFeed.AddPost(post1)
	newsFeed.AddPost(post2)

	// View the news feed
	err := newsFeed.ViewNewsFeed()
	if err != nil {
		fmt.Println(err)
	}

	// Get a specific post by title
	retrievedPost := newsFeed.GetPost("Go Programming")
	if retrievedPost != nil {
		fmt.Println("Retrieved Post:", *retrievedPost)
	} else {
		fmt.Println("Post not found")
	}

	// Print the entire news feed as a string
	fmt.Println(newsFeed.ToString())
}
