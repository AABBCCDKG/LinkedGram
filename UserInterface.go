// A program Team Project
//
// Purdue University -- CS18000 -- Spring 2024 -- Team Project
//
// Author: Purdue CS
// Version: Mar 31, 2024

package main

import "fmt"

// UserInterface defines the methods for user management
type UserInterface interface {
	EditProfile(username, password, email, name string, connections, followers []string, profilePicture interface{})
	ShowInfo() string
	CheckPwd(pwdIn string) bool
	Login(pwdIn string) bool
	IsOnline() bool
}

// User struct implements UserInterface
type User struct {
	Username       string
	Password       string
	Email          string
	Name           string
	Connections    []string
	Followers      []string
	ProfilePicture interface{} // Placeholder for image data
	Online         bool
}

// EditProfile edits the user's profile information
func (u *User) EditProfile(username, password, email, name string, connections, followers []string, profilePicture interface{}) {
	u.Username = username
	u.Password = password
	u.Email = email
	u.Name = name
	u.Connections = connections
	u.Followers = followers
	u.ProfilePicture = profilePicture
}

// ShowInfo returns a string containing user information
func (u *User) ShowInfo() string {
	return fmt.Sprintf("User: %s\nEmail: %s\nName: %s\n", u.Username, u.Email, u.Name)
}

// CheckPwd checks if the provided password matches the user's password
func (u *User) CheckPwd(pwdIn string) bool {
	return u.Password == pwdIn
}

// Login logs the user in if the password is correct
func (u *User) Login(pwdIn string) bool {
	if u.CheckPwd(pwdIn) {
		u.Online = true
		return true
	}
	return false
}

// IsOnline checks if the user is online
func (u *User) IsOnline() bool {
	return u.Online
}

func main() {
	// Example usage:
	user := &User{
		Username: "Alice",
		Password: "password123",
		Email:    "alice@example.com",
		Name:     "Alice Wonderland",
		Online:   false,
	}

	// Editing profile
	user.EditProfile("Alice", "newpassword", "alice@example.com", "Alice Wonderland", []string{}, []string{}, nil)

	// Showing user info
	fmt.Println(user.ShowInfo())

	// Checking password and login
	if user.Login("newpassword") {
		fmt.Println("Login successful!")
	} else {
		fmt.Println("Login failed!")
	}

	// Checking if user is online
	fmt.Printf("Is user online? %v\n", user.IsOnline())
}
