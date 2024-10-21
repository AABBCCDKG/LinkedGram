// A program Team Project
//
// Purdue University -- CS18000 -- Spring 2024 -- Team Project
//
// Author: Purdue CS
// Version: Mar 31, 2024

package main

import (
	"testing"
)

// TestAddComment tests the AddComment method of Post
func TestAddComment(t *testing.T) {
	post, _ := NewPost("title", "content", "author", "url", 0, 0, false, "")
	err := post.AddComment("author", "message")
	if err != nil {
		t.Fatalf("Failed to add comment: %s", err)
	}
	if len(post.GetComments()) != 1 {
		t.Errorf("Expected 1 comment, got %d", len(post.GetComments()))
	}
}

// TestUpvote tests the Upvote method of Post
func TestUpvote(t *testing.T) {
	post, _ := NewPost("title", "content", "author", "url", 0, 0, false, "")
	post.Upvote()
	if post.GetUpvotes() != 1 {
		t.Errorf("Expected 1 upvote, got %d", post.GetUpvotes())
	}
}

// TestDownvote tests the Downvote method of Post
func TestDownvote(t *testing.T) {
	post, _ := NewPost("title", "content", "author", "url", 0, 0, false, "")
	post.Downvote()
	if post.GetDownvotes() != 1 {
		t.Errorf("Expected 1 downvote, got %d", post.GetDownvotes())
	}
}

// TestHide tests the Hide method of Post
func TestHide(t *testing.T) {
	post, _ := NewPost("title", "content", "author", "url", 0, 0, false, "")
	post.Hide()
	if !post.IsHidden() {
		t.Errorf("Expected post to be hidden")
	}
}

// TestShow tests the Show method of Post
func TestShow(t *testing.T) {
	post, _ := NewPost("title", "content", "author", "url", 0, 0, true, "")
	post.Show()
	if post.IsHidden() {
		t.Errorf("Expected post to be visible")
	}
}

// TestSetTitle tests the SetTitle method of Post
func TestSetTitle(t *testing.T) {
	post, _ := NewPost("title", "content", "author", "url", 0, 0, false, "")
	post.SetTitle("new title")
	if post.GetTitle() != "new title" {
		t.Errorf("Expected title to be 'new title', got '%s'", post.GetTitle())
	}
}

// TestSetContent tests the SetContent method of Post
func TestSetContent(t *testing.T) {
	post, _ := NewPost("title", "content", "author", "url", 0, 0, false, "")
	post.SetContent("new content")
	if post.GetContent() != "new content" {
		t.Errorf("Expected content to be 'new content', got '%s'", post.GetContent())
	}
}

// TestSetAuthor tests the SetAuthor method of Post
func TestSetAuthor(t *testing.T) {
	post, _ := NewPost("title", "content", "author", "url", 0, 0, false, "")
	post.SetAuthor("new author")
	if post.GetAuthor() != "new author" {
		t.Errorf("Expected author to be 'new author', got '%s'", post.GetAuthor())
	}
}

// TestSetImageURL tests the SetImageURL method of Post
func TestSetImageURL(t *testing.T) {
	post, _ := NewPost("title", "content", "author", "url", 0, 0, false, "")
	post.SetImageURL("new url")
	if post.GetImageURL() != "new url" {
		t.Errorf("Expected imageURL to be 'new url', got '%s'", post.GetImageURL())
	}
}

// TestSetUpvotes tests the SetUpvotes method of Post
func TestSetUpvotes(t *testing.T) {
	post, _ := NewPost("title", "content", "author", "url", 0, 0, false, "")
	post.SetUpvotes(10)
	if post.GetUpvotes() != 10 {
		t.Errorf("Expected 10 upvotes, got %d", post.GetUpvotes())
	}
}

// TestSetDownvotes tests the SetDownvotes method of Post
func TestSetDownvotes(t *testing.T) {
	post, _ := NewPost("title", "content", "author", "url", 0, 0, false, "")
	post.SetDownvotes(10)
	if post.GetDownvotes() != 10 {
		t.Errorf("Expected 10 downvotes, got %d", post.GetDownvotes())
	}
}

// TestSetHidden tests the SetHidden method of Post
func TestSetHidden(t *testing.T) {
	post, _ := NewPost("title", "content", "author", "url", 0, 0, false, "")
	post.SetHidden(true)
	if !post.IsHidden() {
		t.Errorf("Expected post to be hidden")
	}
}
