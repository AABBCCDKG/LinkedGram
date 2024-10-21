// A program Team Project
//
// Purdue University -- CS18000 -- Spring 2024 -- Team Project
//
// Author: Purdue CS
// Version: Mar 31, 2024

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Post struct simulates the Post class
type Post struct {
	Title     string
	Content   string
	Author    string
	ImageURL  string
	Upvotes   int
	Downvotes int
	Hidden    bool
}

// NewsFeedManager handles post management
type NewsFeedManager struct {
	PostsFile string
}

// NewNewsFeedManager creates a new instance of NewsFeedManager
func NewNewsFeedManager(postsFile string) *NewsFeedManager {
	return &NewsFeedManager{
		PostsFile: postsFile,
	}
}

// GetAllPosts reads all posts from the file
func (manager *NewsFeedManager) GetAllPosts() ([]Post, error) {
	posts := []Post{}

	file, err := os.Open(manager.PostsFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		post, err := stringToPost(line)
		if err != nil {
			fmt.Println("Error parsing post:", err)
			continue
		}
		posts = append(posts, post)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// stringToPost converts a string to a Post object
func stringToPost(line string) (Post, error) {
	parts := strings.Split(line, ";")
	if len(parts) != 7 {
		return Post{}, fmt.Errorf("incomplete post data")
	}

	upvotes, err := strconv.Atoi(parts[4])
	if err != nil {
		return Post{}, err
	}

	downvotes, err := strconv.Atoi(parts[5])
	if err != nil {
		return Post{}, err
	}

	hidden, err := strconv.ParseBool(parts[6])
	if err != nil {
		return Post{}, err
	}

	return Post{
		Title:     parts[0],
		Content:   parts[1],
		Author:    parts[2],
		ImageURL:  parts[3],
		Upvotes:   upvotes,
		Downvotes: downvotes,
		Hidden:    hidden,
	}, nil
}

func main() {
	// Create a new instance of NewsFeedManager
	manager := NewNewsFeedManager("posts.txt")

	// Get all posts from the file
	posts, err := manager.GetAllPosts()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print out the posts
	for _, post := range posts {
		fmt.Println(post)
	}
}
