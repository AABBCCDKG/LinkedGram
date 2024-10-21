package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ClientGUI struct {
	usernameField *widget.Entry
	passwordField *widget.Entry
	socket        net.Conn
	out           *bufio.Writer
	in            *bufio.Reader
	loggedInUser  string
	frame         fyne.Window
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Social Media Platform Client")
	myWindow.Resize(fyne.NewSize(500, 400))

	client := &ClientGUI{
		usernameField: widget.NewEntry(),
		passwordField: widget.NewPasswordEntry(),
		frame:         myWindow,
	}

	client.initializeUI()
	client.connectToServer()

	myWindow.ShowAndRun()
}

func (c *ClientGUI) connectToServer() {
	var err error
	c.socket, err = net.Dial("tcp", "localhost:8081")
	if err != nil {
		c.showDialog("Cannot connect to server: " + err.Error())
		return
	}
	c.out = bufio.NewWriter(c.socket)
	c.in = bufio.NewReader(c.socket)
	fmt.Println("Connected to server.")
}

func (c *ClientGUI) initializeUI() {
	// Create login panel
	usernameLabel := widget.NewLabel("Username:")
	passwordLabel := widget.NewLabel("Password:")

	loginButton := widget.NewButton("Login", func() {
		c.loginAction()
	})

	registerButton := widget.NewButton("Register", func() {
		c.registerAction()
	})

	form := container.NewVBox(
		usernameLabel, c.usernameField,
		passwordLabel, c.passwordField,
		loginButton, registerButton,
	)

	c.frame.SetContent(form)
}

func (c *ClientGUI) loginAction() {
	username := c.usernameField.Text
	password := c.passwordField.Text
	_, err := fmt.Fprintf(c.out, "LOGIN\n%s\n%s\n", username, password)
	if err != nil {
		c.showDialog("Error sending login request: " + err.Error())
		return
	}
	c.out.Flush()

	response, err := c.in.ReadString('\n')
	if err != nil {
		c.showDialog("Error reading server response: " + err.Error())
		return
	}

	if strings.TrimSpace(response) == "true" {
		c.loggedInUser = username
		c.showDialog("Login successful!")
		c.showSocialMediaMenu()
	} else {
		c.showDialog("Login failed!")
	}
}

func (c *ClientGUI) registerAction() {
	username := c.usernameField.Text
	password := c.passwordField.Text
	_, err := fmt.Fprintf(c.out, "REGISTER\n%s\n%s\n", username, password)
	if err != nil {
		c.showDialog("Error sending registration request: " + err.Error())
		return
	}
	c.out.Flush()

	response, err := c.in.ReadString('\n')
	if err != nil {
		c.showDialog("Error reading server response: " + err.Error())
		return
	}
	c.showDialog(strings.TrimSpace(response))
}

func (c *ClientGUI) showSocialMediaMenu() {
	viewFeedButton := widget.NewButton("View News Feed", func() {
		// handle viewing feed
		c.showDialog("Viewing news feed...")
	})

	createPostButton := widget.NewButton("Create Post", func() {
		// handle post creation
		c.showDialog("Creating post...")
	})

	logoutButton := widget.NewButton("Logout", func() {
		c.logout()
	})

	menu := container.NewVBox(viewFeedButton, createPostButton, logoutButton)
	c.frame.SetContent(menu)
}

func (c *ClientGUI) logout() {
	_, err := fmt.Fprintf(c.out, "LOGOUT\n")
	if err != nil {
		c.showDialog("Error logging out: " + err.Error())
		return
	}
	c.out.Flush()

	response, err := c.in.ReadString('\n')
	if err != nil {
		c.showDialog("Error reading server response: " + err.Error())
		return
	}
	if strings.TrimSpace(response) == "Logout successful!" {
		c.showDialog("Logged out successfully!")
		c.initializeUI()
	} else {
		c.showDialog("Failed to logout: " + response)
	}
}

func (c *ClientGUI) showDialog(message string) {
	dialog := widget.NewLabel(message)
	dialogWindow := c.frame
	dialogWindow.SetContent(dialog)
	dialogWindow.Show()
}
