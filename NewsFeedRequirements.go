// A program Team Project
//
// Purdue University -- CS18000 -- Spring 2024 -- Team Project
//
// Author: Purdue CS
// Version: Mar 31, 2024

package main

import "fmt"

// Post struct to represent a post
type Post struct {
	Title     string
	Content   string
	ImageURL  string
	Author    string
	Hidden    bool
	Upvotes   int
	Downvotes int
}

// NewsFeedRequirements defines the methods required by the NewsFeed
type NewsFeedRequirements interface {
	AddPost(post Post)
	RemovePost(post Post)
	HidePost(post Post)
	ShowPost(post Post)
	DisplayNewsFeed()
	DisplayPost(post Post)
	EditPost(post *Post, title, content, imageURL string)
}

// NewsFeed struct that implements NewsFeedRequirements interface
type NewsFeed struct {
	Posts []Post
}

// AddPost adds a post to the NewsFeed
func (n *NewsFeed) AddPost(post Post) {
	n.Posts = append(n.Posts, post)
}

// RemovePost removes a post from the NewsFeed
func (n *NewsFeed) RemovePost(post Post) {
	for i, p := range n.Posts {
		if p.Title == post.Title {
			n.Posts = append(n.Posts[:i], n.Posts[i+1:]...)
			break
		}
	}
}

// HidePost hides a post in the NewsFeed
func (n *NewsFeed) HidePost(post Post) {
	for i, p := range n.Posts {
		if p.Title == post.Title {
			n.Posts[i].Hidden = true
			break
		}
	}
}

// ShowPost makes a post visible in the NewsFeed
func (n *NewsFeed) ShowPost(post Post) {
	for i, p := range n.Posts {
		if p.Title == post.Title {
			n.Posts[i].Hidden = false
			break
		}
	}
}

// DisplayNewsFeed displays all posts in the NewsFeed
func (n *NewsFeed) DisplayNewsFeed() {
	for _, post := range n.Posts {
		n.DisplayPost(post)
	}
}

// DisplayPost displays a single post
func (n *NewsFeed) DisplayPost(post Post) {
	if post.Hidden {
		fmt.Println("This post is hidden.")
	} else {
		fmt.Printf("Title: %s\nContent: %s\nImageURL: %s\nAuthor: %s\nUpvotes: %d\nDownvotes: %d\n", post.Title, post.Content, post.ImageURL, post.Author, post.Upvotes, post.Downvotes)
	}
}

// EditPost edits an existing post in the NewsFeed
func (n *NewsFeed) EditPost(post *Post, title, content, imageURL string) {
	post.Title = title
	post.Content = content
	post.ImageURL = imageURL
}

func main() {
	// Example usage
	newsFeed := &NewsFeed{}

	post1 := Post{Title: "First Post", Content: "This is the first post.", ImageURL: "image1.jpg", Author: "Author1"}
	post2 := Post{Title: "Second Post", Content: "This is the second post.", ImageURL: "image2.jpg", Author: "Author2"}

	newsFeed.AddPost(post1)
	newsFeed.AddPost(post2)

	newsFeed.DisplayNewsFeed()

	// Hide a post
	newsFeed.HidePost(post1)
	newsFeed.DisplayNewsFeed()

	// Edit a post
	newsFeed.EditPost(&post2, "Updated Title", "Updated content", "updatedImage.jpg")
	newsFeed.DisplayNewsFeed()
}
