// A program Team Project
//
// Purdue University -- CS18000 -- Spring 2024 -- Team Project
//
// Author: Purdue CS
// Version: Mar 31, 2024

package main

// Profile struct is assumed to be defined elsewhere
type Profile struct {
	About string
	// other profile fields...
}

// User struct to represent a user
type User struct {
	Username  string
	Password  string
	Following []*User // Users that this user follows
	Profile   *Profile
}

// NewUser creates a new User instance
func NewUser(username, password string) *User {
	return &User{
		Username:  username,
		Password:  password,
		Following: []*User{},
	}
}

// GetUsername returns the username of the user
func (u *User) GetUsername() string {
	return u.Username
}

// GetPassword returns the password of the user
func (u *User) GetPassword() string {
	return u.Password
}

// GetFollowing returns the list of users that this user is following
func (u *User) GetFollowing() []*User {
	return u.Following
}

// AddFollowing adds another user to the following list
func (u *User) AddFollowing(user *User) {
	u.Following = append(u.Following, user)
}

// RemoveFollowing removes a user from the following list
func (u *User) RemoveFollowing(user *User) {
	for i, followingUser := range u.Following {
		if followingUser.Username == user.Username {
			u.Following = append(u.Following[:i], u.Following[i+1:]...)
			break
		}
	}
}

// GetProfile returns the profile of the user
func (u *User) GetProfile() *Profile {
	return u.Profile
}

// SetProfile sets the profile of the user
func (u *User) SetProfile(profile *Profile) {
	u.Profile = profile
}

// SetUsername changes the username of the user
func (u *User) SetUsername(newUsername string) {
	u.Username = newUsername
}

// AddFollower allows a user to follow another user
func (u *User) AddFollower(follower *User) {
	follower.AddFollowing(u)
}

func main() {
	// Example usage
	user1 := NewUser("Alice", "password123")
	user2 := NewUser("Bob", "password456")

	// Alice follows Bob
	user1.AddFollowing(user2)

	// Display following list for Alice
	for _, user := range user1.GetFollowing() {
		println("Alice follows:", user.GetUsername())
	}

	// Bob follows Alice back
	user2.AddFollower(user1)

	// Display following list for Bob
	for _, user := range user2.GetFollowing() {
		println("Bob follows:", user.GetUsername())
	}
}
