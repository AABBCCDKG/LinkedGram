package main

import (
	"fmt"
)

// PostIncompleteError represents an error where a post is incomplete
type PostIncompleteError struct {
	Message string
}

// Error implements the error interface
func (e *PostIncompleteError) Error() string {
	return fmt.Sprintf("post incomplete: %s", e.Message)
}

// NewPostIncompleteError creates a new PostIncompleteError with the given message
func NewPostIncompleteError(message string) error {
	return &PostIncompleteError{
		Message: message,
	}
}

// 示范如何在函数中使用这个自定义错误类型
func parsePost(data string) (*Post, error) {
	// 假设这里有一些解析逻辑
	if len(data) == 0 {
		return nil, NewPostIncompleteError("post data is empty")
	}

	// 假设解析成功
	post := &Post{
		Title:       "Example Post",
		Content:     "This is an example post.",
		Author:      "Author Name",
		ImageURL:    "http://example.com/image.jpg",
		Upvotes:     10,
		Downvotes:   2,
		IsHidden:    false,
		Description: "This is a description of the post.",
	}
	return post, nil
}

func main() {
	postData := ""
	post, err := parsePost(postData)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Parsed Post:", post)
}