// A program Team Project
//
// Purdue University -- CS18000 -- Spring 2024 -- Team Project
//
// Author: Purdue CS
// Version: Mar 31, 2024

package main

import (
	"testing"
)

// TestGetUsername tests the GetUsername method
func TestGetUsername(t *testing.T) {
	user := NewUser("username", "password")
	if user.GetUsername() != "username" {
		t.Errorf("Expected 'username', got '%s'", user.GetUsername())
	}
}

// TestGetPassword tests the GetPassword method
func TestGetPassword(t *testing.T) {
	user := NewUser("username", "password")
	if user.GetPassword() != "password" {
		t.Errorf("Expected 'password', got '%s'", user.GetPassword())
	}
}

// TestGetFollowing tests the GetFollowing method
func TestGetFollowing(t *testing.T) {
	user1 := NewUser("username1", "password1")
	user2 := NewUser("username2", "password2")
	user1.AddFollowing(user2)
	if !contains(user1.GetFollowing(), user2) {
		t.Errorf("Expected user1 to follow user2")
	}
}

// TestAddFollowing tests the AddFollowing method
func TestAddFollowing(t *testing.T) {
	user1 := NewUser("username1", "password1")
	user2 := NewUser("username2", "password2")
	user1.AddFollowing(user2)
	if !contains(user1.GetFollowing(), user2) {
		t.Errorf("Expected user1 to follow user2")
	}
}

// TestSetUsername tests the SetUsername method
func TestSetUsername(t *testing.T) {
	user := NewUser("username", "password")
	user.SetUsername("newUsername")
	if user.GetUsername() != "newUsername" {
		t.Errorf("Expected 'newUsername', got '%s'", user.GetUsername())
	}
}

// TestRemoveFollowing tests the RemoveFollowing method
func TestRemoveFollowing(t *testing.T) {
	user1 := NewUser("username1", "password1")
	user2 := NewUser("username2", "password2")
	user1.AddFollowing(user2)
	user1.RemoveFollowing(user2)
	if contains(user1.GetFollowing(), user2) {
		t.Errorf("Expected user1 to no longer follow user2")
	}
}

// TestGetProfile tests the GetProfile method
func TestGetProfile(t *testing.T) {
	user := NewUser("username", "password")
	profile := &Profile{
		About:      "about",
		Education:  "education",
		Awards:     "awards",
		Skills:     "skills",
		Status:     "status",
	}
	user.SetProfile(profile)
	if user.GetProfile() != profile {
		t.Errorf("Expected profile, got different profile")
	}
}

// TestSetProfile tests the SetProfile method
func TestSetProfile(t *testing.T) {
	user := NewUser("username", "password")
	profile := &Profile{
		About:      "about",
		Education:  "education",
		Awards:     "awards",
		Skills:     "skills",
		Status:     "status",
	}
	user.SetProfile(profile)
	if user.GetProfile() != profile {
		t.Errorf("Expected profile, got different profile")
	}
}

// TestAddFollower tests the AddFollower method
func TestAddFollower(t *testing.T) {
	user1 := NewUser("username1", "password1")
	user2 := NewUser("username2", "password2")
	user1.AddFollower(user2)
	if !contains(user2.GetFollowing(), user1) {
		t.Errorf("Expected user2 to follow user1")
	}
}

// Utility function to check if a user is in a slice of users
func contains(users []*User, user *User) bool {
	for _, u := range users {
		if u == user {
			return true
		}
	}
	return false
}
