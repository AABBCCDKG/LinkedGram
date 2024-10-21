package main

import (
    "testing"
)

type Post struct {
    Title   string
    Content string
    Author  string
    URL     string
}

type Comment struct {
    Post      *Post
    Content   string
    Author    string
    Hidden    bool
    Upvotes   int
    Downvotes int
}

func (c *Comment) upvote() {
    c.Upvotes++
}

func (c *Comment) downvote() {
    c.Downvotes++
}

func (c *Comment) hide() {
    c.Hidden = true
}

func (c *Comment) show() {
    c.Hidden = false
}

func (c *Comment) setContent(content string) {
    c.Content = content
}

func (c *Comment) setAuthor(author string) {
    c.Author = author
}

func (c *Comment) setHidden(hidden bool) {
    c.Hidden = hidden
}

func (c *Comment) setUpvotes(upvotes int) {
    c.Upvotes = upvotes
}

func (c *Comment) setDownvotes(downvotes int) {
    c.Downvotes = downvotes
}

func (c *Comment) setPost(post *Post) {
    c.Post = post
}

func TestUpvote(t *testing.T) {
    comment := &Comment{
        Post:    &Post{"title", "content", "author", "url"},
        Content: "content",
        Author:  "author",
    }
    comment.upvote()
    if comment.Upvotes != 1 {
        t.Errorf("Expected 1 upvote, got %d", comment.Upvotes)
    }
}

func TestDownvote(t *testing.T) {
    comment := &Comment{
        Post:    &Post{"title", "content", "author", "url"},
        Content: "content",
        Author:  "author",
    }
    comment.downvote()
    if comment.Downvotes != 1 {
        t.Errorf("Expected 1 downvote, got %d", comment.Downvotes)
    }
}

func TestHide(t *testing.T) {
    comment := &Comment{
        Post:    &Post{"title", "content", "author", "url"},
        Content: "content",
        Author:  "author",
    }
    comment.hide()
    if !comment.Hidden {
        t.Errorf("Expected comment to be hidden")
    }
}

func TestShow(t *testing.T) {
    comment := &Comment{
        Post:    &Post{"title", "content", "author", "url"},
        Content: "content",
        Author:  "author",
        Hidden:  true,
    }
    comment.show()
    if comment.Hidden {
        t.Errorf("Expected comment to be shown")
    }
}

func TestSetContent(t *testing.T) {
    comment := &Comment{
        Post:    &Post{"title", "content", "author", "url"},
        Content: "content",
        Author:  "author",
    }
    comment.setContent("new content")
    if comment.Content != "new content" {
        t.Errorf("Expected content to be 'new content', got %s", comment.Content)
    }
}

func TestSetAuthor(t *testing.T) {
    comment := &Comment{
        Post:    &Post{"title", "content", "author", "url"},
        Content: "content",
        Author:  "author",
    }
    comment.setAuthor("new author")
    if comment.Author != "new author" {
        t.Errorf("Expected author to be 'new author', got %s", comment.Author)
    }
}

func TestSetHidden(t *testing.T) {
    comment := &Comment{
        Post:    &Post{"title", "content", "author", "url"},
        Content: "content",
        Author:  "author",
    }
    comment.setHidden(true)
    if !comment.Hidden {
        t.Errorf("Expected comment to be hidden")
    }
}

func TestSetUpvotes(t *testing.T) {
    comment := &Comment{
        Post:    &Post{"title", "content", "author", "url"},
        Content: "content",
        Author:  "author",
    }
    comment.setUpvotes(10)
    if comment.Upvotes != 10 {
        t.Errorf("Expected 10 upvotes, got %d", comment.Upvotes)
    }
}

func TestSetDownvotes(t *testing.T) {
    comment := &Comment{
        Post:    &Post{"title", "content", "author", "url"},
        Content: "content",
        Author:  "author",
    }
    comment.setDownvotes(10)
    if comment.Downvotes != 10 {
        t.Errorf("Expected 10 downvotes, got %d", comment.Downvotes)
    }
}

func TestSetPost(t *testing.T) {
    comment := &Comment{
        Post:    &Post{"title", "content", "author", "url"},
        Content: "content",
        Author:  "author",
    }
    newPost := &Post{"new title", "new content", "new author", "new url"}
    comment.setPost(newPost)
    if comment.Post != newPost {
        t.Errorf("Expected post to be %v, got %v", newPost, comment.Post)
    }
}