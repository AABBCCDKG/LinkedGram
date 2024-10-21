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
	"strings"
	"sync"
)

const fileName = "user_credentials.txt"

// credentials stores the username-password pairs in memory
var credentials = make(map[string]string)

// mutex to handle concurrent access to the credentials map
var mutex = &sync.Mutex{}

// Initialize by loading credentials from file
func init() {
	err := loadCredentialsFromFile()
	if err != nil {
		fmt.Println("Error loading user credentials from file:", err)
	}
}

// IsValidCredentials checks if the given username and password are correct
func IsValidCredentials(username, password string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	storedPassword, exists := credentials[username]
	return exists && storedPassword == password
}

// AddCredentials adds a new username and password, and saves it to the file
func AddCredentials(username, password string) error {
	mutex.Lock()
	defer mutex.Unlock()
	credentials[username] = password
	return saveCredentialsToFile()
}

// loadCredentialsFromFile loads the credentials from the file
func loadCredentialsFromFile() error {
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			// File does not exist yet, that's okay
			return nil
		}
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			credentials[parts[0]] = parts[1]
		}
	}
	return scanner.Err()
}

// saveCredentialsToFile saves the credentials to the file
func saveCredentialsToFile() error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for username, password := range credentials {
		_, err := writer.WriteString(username + ":" + password + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func main() {
	// Example usage:
	fmt.Println("Adding credentials...")
	err := AddCredentials("Alice", "password123")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Checking credentials for Alice...")
	isValid := IsValidCredentials("Alice", "password123")
	fmt.Println("Are credentials valid?", isValid)
}
