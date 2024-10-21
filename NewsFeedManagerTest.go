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

// TestGetAllPosts tests the GetAllPosts function of NewsFeedManager
func TestGetAllPosts(t *testing.T) {
	// 创建一个临时的 posts.txt 文件用于测试
	content := "Title1;Content1;Author1;ImageURL1;5;3;false\n" +
		"Title2;Content2;Author2;ImageURL2;10;1;true\n"
	file, err := os.CreateTemp("", "posts.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %s", err)
	}
	defer os.Remove(file.Name()) // 在测试结束后删除临时文件

	// 写入内容到文件
	if _, err := file.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %s", err)
	}

	// 创建 NewsFeedManager 实例
	manager := NewNewsFeedManager(file.Name())

	// 调用 GetAllPosts 方法
	posts, err := manager.GetAllPosts()
	if err != nil {
		t.Fatalf("GetAllPosts failed: %s", err)
	}

	// 假设 posts.txt 文件中至少有一个帖子
	if len(posts) == 0 {
		t.Fatalf("Expected at least one post, got none")
	}

	// 检查第一个帖子的属性
	firstPost := posts[0]
	if firstPost.Title == "" {
		t.Error("Expected non-empty title")
	}
	if firstPost.Content == "" {
		t.Error("Expected non-empty content")
	}
	if firstPost.Author == "" {
		t.Error("Expected non-empty author")
	}
	if firstPost.ImageURL == "" {
		t.Error("Expected non-empty image URL")
	}
	if firstPost.Upvotes == 0 {
		t.Error("Expected upvotes to be non-zero")
	}
	if firstPost.Downvotes == 0 {
		t.Error("Expected downvotes to be non-zero")
	}
	if firstPost.Hidden {
		t.Error("Expected hidden to be false")
	}
}
