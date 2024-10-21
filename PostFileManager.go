package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 假设Post的定义如下：
type Post struct {
	Title       string
	Content     string
	Author      string
	ImageURL    string
	Upvotes     int
	Downvotes   int
	IsHidden    bool
	Description string
}

var postsFile = "posts.txt"

// SavePost saves a post to the file
func SavePost(post *Post) {
	file, err := os.OpenFile(postsFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(postToString(post) + "\n\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	writer.Flush()
	fmt.Println("Post saved to file:", post.Title)
}

// GetAllPosts reads all posts from the file
func GetAllPosts() []*Post {
	file, err := os.Open(postsFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var posts []*Post
	line, err := reader.ReadString('\n')
	for err == nil {
		post, err := stringToPost(line)
		if err != nil {
			fmt.Println("Error parsing post:", err)
			return nil
		}
		posts = append(posts, post)
		line, err = reader.ReadString('\n')
	}
	if err != nil && err != io.EOF {
		fmt.Println("Error reading file:", err)
		return nil
	}
	return posts
}

func postToString(post *Post) string {
	return fmt.Sprintf("%s;%s;%s;%s;%d;%d;%t;%s",
		post.Title, post.Content, post.Author, post.ImageURL, post.Upvotes, post.Downvotes, post.IsHidden, post.Description)
}

func stringToPost(line string) (*Post, error) {
	parts := strings.Split(line, ";")
	if len(parts) != 8 {
		return nil, fmt.Errorf("invalid post format")
	}
	upvotes, err := strconv.Atoi(parts[4])
	if err != nil {
		return nil, fmt.Errorf("invalid upvotes: %v", err)
	}
	downvotes, err := strconv.Atoi(parts[5])
	if err != nil {
		return nil, fmt.Errorf("invalid downvotes: %v", err)
	}
	isHidden, err := strconv.ParseBool(parts[6])
	if err != nil {
		return nil, fmt.Errorf("invalid isHidden: %v", err)
	}
	return &Post{
		Title:       parts[0],
		Content:     parts[1],
		Author:      parts[2],
		ImageURL:    parts[3],
		Upvotes:     upvotes,
		Downvotes:   downvotes,
		IsHidden:    isHidden,
		Description: parts[7],
	}, nil
}

func main() {
	// Example usage
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
	SavePost(post)

	allPosts := GetAllPosts()
	for _, post := range allPosts {
		fmt.Println(postToString(post))
	}
}