package main

import "fmt"

// ProfileInterface defines the interface for profile operations
type ProfileInterface interface {
	AddUser(user *User)
	DeleteUser(user *User)
	AddExperience(experience string)
	AddFollower(follower string, user *User)
	RemoveFollower(follower string, user *User)
	AddConnection(connection string, user *User)
	RemoveConnection(connection string, user *User)
}

// User represents a user entity
type User struct {
	Name   string
	Email  string
	Experience []string
	Followers map[string]*User
	Connections map[string]*User
}

// NewUser creates a new user
func NewUser(name, email string) *User {
	return &User{
		Name:       name,
		Email:      email,
		Experience: []string{},
		Followers:  make(map[string]*User),
		Connections: make(map[string]*User),
	}
}

// Profile represents a user's profile
type Profile struct {
	User *User
}

// NewProfile creates a new profile
func NewProfile(name, email string) *Profile {
	return &Profile{
		User: NewUser(name, email),
	}
}

// AddUser adds a user to the profile
func (p *Profile) AddUser(user *User) {
	p.User = user
}

// DeleteUser deletes a user from the profile
func (p *Profile) DeleteUser(user *User) {
	// In this example, we simply set the user to nil
	p.User = nil
}

// AddExperience adds an experience to the user's profile
func (p *Profile) AddExperience(experience string) {
	p.User.Experience = append(p.User.Experience, experience)
}

// AddFollower adds a follower to the user's profile
func (p *Profile) AddFollower(follower string, user *User) {
	p.User.Followers[follower] = user
}

// RemoveFollower removes a follower from the user's profile
func (p *Profile) RemoveFollower(follower string, user *User) {
	delete(p.User.Followers, follower)
}

// AddConnection adds a connection to the user's profile
func (p *Profile) AddConnection(connection string, user *User) {
	p.User.Connections[connection] = user
}

// RemoveConnection removes a connection from the user's profile
func (p *Profile) RemoveConnection(connection string, user *User) {
	delete(p.User.Connections, connection)
}

func main() {
	prof := NewProfile("John Doe", "john.doe@example.com")

	// Implement the ProfileInterface methods
	prof.AddUser(NewUser("Jane Doe", "jane.doe@example.com"))
	prof.AddExperience("Senior Software Developer at XYZ Corp")
	prof.AddFollower("alice123", NewUser("Alice Smith", "alice.smith@example.com"))
	prof.AddConnection("bob123", NewUser("Bob Johnson", "bob.johnson@example.com"))

	fmt.Printf("User: %+v\n", prof.User)
}