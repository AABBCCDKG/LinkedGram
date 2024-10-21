package main

import (
	"testing"
)

// 假设Post和NewsFeed的定义如下：
type Post struct {
	Title   string
	Content string
	Author  string
	IsDraft bool
	URL     string
	Likes   int
	Shares  int
}

type NewsFeed struct {
	posts map[string]*Post
}

func NewNewsFeed() *NewsFeed {
	return &NewsFeed{posts: make(map[string]*Post)}
}

func (nf *NewsFeed) AddPost(post *Post) {
	nf.posts[post.Title] = post
}

func (nf *NewsFeed) RemovePost(post *Post) {
	delete(nf.posts, post.Title)
}

func (nf *NewsFeed) GetPost(title string) *Post {
	return nf.posts[title]
}

func (nf *NewsFeed) String() string {
	var result string
	for _, post := range nf.posts {
		result += post.String() + "\n"
	}
	return result
}

func (p *Post) String() string {
	return p.Title + ": " + p.Content + "\n"
}

// 测试方法
func TestAddPost(t *testing.T) {
	newsFeed := NewNewsFeed()
	post := &Post{"title", "content", "author", false, "url", 0, 0}
	newsFeed.AddPost(post)
	if newsFeed.getPost("title") != post {
		t.Errorf("Expected post %v, got %v", post, newsFeed.getPost("title"))
	}
}

func TestRemovePost(t *testing.T) {
	newsFeed := NewNewsFeed()
	post := &Post{"title", "content", "author", false, "url", 0, 0}
	newsFeed.AddPost(post)
	newsFeed.RemovePost(post)
	if post, exists := newsFeed.posts["title"]; exists {
		t.Errorf("Expected no post, got %v", post)
	}
}

func TestGetPost(t *testing.T) {
	newsFeed := NewNewsFeed()
	post := &Post{"title", "content", "author", false, "url", 0, 0}
	newsFeed.AddPost(post)
	if newsFeed.getPost("title") != post {
		t.Errorf("Expected post %v, got %v", post, newsFeed.getPost("title"))
	}
}

func TestToString(t *testing.T) {
	newsFeed := NewNewsFeed()
	post := &Post{"title", "content", "author", false, "url", 0, 0}
	newsFeed.AddPost(post)
	expected := post.String() + "\n"
	if newsFeed.String() != expected {
		t.Errorf("Expected string %v, got %v", expected, newsFeed.String())
	}
}