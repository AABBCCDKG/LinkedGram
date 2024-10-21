// A program Team Project
//
// Purdue University -- CS18000 -- Spring 2024 -- Team Project
//
// Author: Purdue CS
// Version: Mar 31, 2024

package main

import (
	"os"
	"testing"
)

// PostFileManager handles file operations for posts
type PostFileManager struct {
	PostsFile string
}

// SavePost saves a post to the posts file
func (manager *PostFileManager) SavePost(post *Post) error {
	file, err := os.OpenFile(manager.PostsFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(post.ToString() + "\n")
	if err != nil {
		return err
	}

	return nil
}

// GetAllPosts loads all posts from the posts file
func (manager *PostFileManager) GetAllPosts() ([]Post, error) {
	posts := []Post{}
	file, err := os.Open(manager.PostsFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Load posts (assuming some logic to deserialize)
	// For simplicity, we assume all posts are stored in the same format as their ToString output

	return posts, nil
}

// TestSavePost tests the SavePost method of PostFileManager
func TestSavePost(t *testing.T) {
	// Create a temporary file for testing
	file, err := os.CreateTemp("", "posts.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %s", err)
	}
	defer os.Remove(file.Name())

	manager := &PostFileManager{PostsFile: file.Name()}

	// Create a new post
	post, err := NewPost("title", "content", "author", "imageURL", 0, 0, false, "comments")
	if err != nil {
		t.Fatalf("Failed to create post: %s", err)
	}

	// Save the post to the file
	err = manager.SavePost(post)
	if err != nil {
		t.Fatalf("Failed to save post: %s", err)
	}

	// Load posts and check if the saved post is present
	posts, err := manager.GetAllPosts()
	if err != nil {
		t.Fatalf("Failed to get all posts: %s", err)
	}

	found := false
	for _, p := range posts {
		if p.GetTitle() == post.GetTitle() {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Post was not found in saved posts")
	}
}

// TestGetAllPosts tests the GetAllPosts method of PostFileManager
func TestGetAllPosts(t *testing.T) {
	// Create a temporary file for testing
	file, err := os.CreateTemp("", "posts.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %s", err)
	}
	defer os.Remove(file.Name())

	manager := &PostFileManager{PostsFile: file.Name()}

	// Assuming the file contains at least one post (for this test we could add some sample posts)

	posts, err := manager.GetAllPosts()
	if err != nil {
		t.Fatalf("Failed to get all posts: %s", err)
	}

	if len(posts) == 0 {
		t.Fatalf("Expected at least one post, got none")
	}

	// Validate the first post's properties
	firstPost := posts[0]
	if firstPost.GetTitle() == "" {
		t.Error("Expected non-empty title")
	}
	if firstPost.GetContent() == "" {
		t.Error("Expected non-empty content")
	}
	if firstPost.GetAuthor() == "" {
		t.Error("Expected non-empty author")
	}
	if firstPost.GetImageURL() == "" {
		t.Error("Expected non-empty image URL")
	}
	if firstPost.GetUpvotes() == 0 {
		t.Error("Expected upvotes to be non-zero")
	}
	if firstPost.GetDownvotes() == 0 {
		t.Error("Expected downvotes to be non-zero")
	}
	if !firstPost.IsHidden() {
		t.Error("Expected post to be hidden")
	}
}
