package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type ClientHandler struct {
	clientSocket net.Conn
	bufferedReader *bufio.Reader
	printWriter *bufio.Writer
	users []*User
	newsFeeds []*NewsFeed
}

func NewClientHandler(clientSocket net.Conn, users []*User, newsFeeds []*NewsFeed) *ClientHandler {
	return &ClientHandler{
		clientSocket:   clientSocket,
		bufferedReader: bufio.NewReader(clientSocket),
		printWriter:    bufio.NewWriter(clientSocket),
		users:          users,
		newsFeeds:      newsFeeds,
	}
}

func (ch *ClientHandler) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	defer ch.clientSocket.Close()

	for {
		request, err := ch.bufferedReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from client:", err)
			break
		}

		request = strings.TrimSpace(request)
		response := ch.handleRequest(request)
		ch.printWriter.WriteString(response + "\n")
		ch.printWriter.Flush()
	}
}

func (ch *ClientHandler) handleRequest(request string) string {
	requestParts := strings.SplitN(request, " ", 2)
	switch requestParts[0] {
	case "REGISTER":
		return ch.handleRegistration()
	case "LOGIN":
		return ch.handleLogin()
	case "VIEW_NEWS_FEED":
		return ch.handleViewNewsFeed()
	case "CREATE_POST":
		ch.handleCreatePost()
		return "Post created successfully!"
	case "UPVOTE_POST":
		return ch.handleUpvotePost()
	case "DOWNVOTE_POST":
		return ch.handleDownvotePost()
	case "COMMENT_ON_POST":
		return ch.handleCommentOnPost()
	case "SEARCH_USER":
		usernameToSearch, _ := ch.bufferedReader.ReadString('\n')
		return ch.handleSearchUser(strings.TrimSpace(usernameToSearch))
	case "CREATE_PROFILE":
		username, _ := ch.bufferedReader.ReadString('\n')
		return ch.handleCreateProfile(strings.TrimSpace(username))
	case "LOGOUT":
		return "Logout successful!"
	default:
		return "Unrecognized request: " + request
	}
}

func (ch *ClientHandler) handleRegistration() string {
	username, _ := ch.bufferedReader.ReadString('\n')
	password, _ := ch.bufferedReader.ReadString('\n')

	if UserCredentials.isValidCredentials(strings.TrimSpace(username), strings.TrimSpace(password)) {
		return "Username already exists. Please choose a different username."
	} else {
		UserCredentials.addCredentials(username, password)
		newUser := &User{username: username, password: password}
		ch.users = append(ch.users, newUser)
		return "Registration successful!"
	}
}

func (ch *ClientHandler) handleLogin() string {
	username, _ := ch.bufferedReader.ReadString('\n')
	password, _ := ch.bufferedReader.ReadString('\n')

	if UserCredentials.isValidCredentials(strings.TrimSpace(username), strings.TrimSpace(password)) {
		for _, user := range ch.users {
			if user.username == username {
				return "true"
			}
		}
	}
	return "false"
}

func (ch *ClientHandler) handleViewNewsFeed() string {
	var newsFeedBuilder strings.Builder
	posts := PostFileManager.getAllPosts()

	if len(posts) > 0 {
		for _, post := range posts {
			newsFeedBuilder.WriteString(post.String() + "\n\n")
		}
		return newsFeedBuilder.String()
	}
	return "No posts in the news feed."
}

func (ch *ClientHandler) handleCreatePost() {
	title, _ := ch.bufferedReader.ReadString('\n')
	content, _ := ch.bufferedReader.ReadString('\n')
	author, _ := ch.bufferedReader.ReadString('\n')
	imageURL, _ := ch.bufferedReader.ReadString('\n')

	newPost := &Post{
		Title:    strings.TrimSpace(title),
		Content:  strings.TrimSpace(content),
		Author:   strings.TrimSpace(author),
		ImageURL: strings.TrimSpace(imageURL),
	}

	PostFileManager.savePost(newPost)
}

func (ch *ClientHandler) handleUpvotePost() string {
	postTitle, _ := ch.bufferedReader.ReadString('\n')
	postTitle = strings.TrimSpace(postTitle)

	lines, err := os.ReadFile("posts.txt")
	if err != nil {
		return "Error reading posts file."
	}

	linesList := strings.Split(string(lines), "\n")
	for i, line := range linesList {
		if strings.HasPrefix(line, postTitle) {
			parts := strings.Split(line, ";")
			upvotes, _ := strconv.Atoi(parts[4])
			upvotes++
			parts[4] = strconv.Itoa(upvotes)
			linesList[i] = strings.Join(parts, ";")
			os.WriteFile("posts.txt", []byte(strings.Join(linesList, "\n")), 0644)
			return "Post upvoted successfully!"
		}
	}

	return "Post not found!"
}

func (ch *ClientHandler) handleDownvotePost() string {
	postTitle, _ := ch.bufferedReader.ReadString('\n')
	postTitle = strings.TrimSpace(postTitle)

	lines, err := os.ReadFile("posts.txt")
	if err != nil {
		return "Error reading posts file."
	}

	linesList := strings.Split(string(lines), "\n")
	for i, line := range linesList {
		if strings.HasPrefix(line, postTitle) {
			parts := strings.Split(line, ";")
			downvotes, _ := strconv.Atoi(parts[5])
			downvotes++
			parts[5] = strconv.Itoa(downvotes)
			linesList[i] = strings.Join(parts, ";")
			os.WriteFile("posts.txt", []byte(strings.Join(linesList, "\n")), 0644)
			return "Post downvoted successfully!"
		}
	}

	return "Post not found!"
}

func (ch *ClientHandler) handleCommentOnPost() string {
	postTitle, _ := ch.bufferedReader.ReadString('\n')
	commentAuthor, _ := ch.bufferedReader.ReadString('\n')
	commentContent, _ := ch.bufferedReader.ReadString('\n')

	lines, err := os.ReadFile("posts.txt")
	if err != nil {
		return "Error reading posts file."
	}

	linesList := strings.Split(string(lines), "\n")
	for i, line := range linesList {
		if strings.HasPrefix(line, strings.TrimSpace(postTitle)) {
			parts := strings.Split(line, ";")
			comments := parts[len(parts)-1]
			comments += ";" + commentAuthor + ":" + commentContent
			parts[len(parts)-1] = comments
			linesList[i] = strings.Join(parts, ";")
			os.WriteFile("posts.txt", []byte(strings.Join(linesList, "\n")), 0644)
			return "Comment added successfully!"
		}
	}

	return "Post not found!"
}

func (ch *ClientHandler) handleCreateProfile(username string) string {
	about, _ := ch.bufferedReader.ReadString('\n')
	experience, _ := ch.bufferedReader.ReadString('\n')
	education, _ := ch.bufferedReader.ReadString('\n')
	awards, _ := ch.bufferedReader.ReadString('\n')
	skills, _ := ch.bufferedReader.ReadString('\n')
	status, _ := ch.bufferedReader.ReadString('\n')

	profileFileName := fmt.Sprintf("%s_profile.txt", username)
	profileFile, err := os.Create(profileFileName)
	if err != nil {
		return "Error creating profile file."
	}
	defer profileFile.Close()

	profileData := fmt.Sprintf("About: %s\nExperience: %s\nEducation: %s\nAwards: %s\nSkills: %s\nStatus: %s\n",
		strings.TrimSpace(about), strings.TrimSpace(experience), strings.TrimSpace(education),
		strings.TrimSpace(awards), strings.TrimSpace(skills), strings.TrimSpace(status))

	profileFile.WriteString(profileData)
	return "Profile created successfully!"
