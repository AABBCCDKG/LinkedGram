package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

const (
    HOST = "localhost"
    PORT = "8081"
)

var (
    conn         net.Conn
    writer       *bufio.Writer
    reader       *bufio.Reader
    scanner      *bufio.Scanner
    loggedInUser string
)

func main() {
    var err error
    conn, err = net.Dial("tcp", HOST+":"+PORT)
    if err != nil {
        fmt.Printf("Error connecting to server: %v\n", err)
        return
    }
    defer conn.Close()

    writer = bufio.NewWriter(conn)
    reader = bufio.NewReader(conn)
    scanner = bufio.NewScanner(os.Stdin)

    fmt.Println("Connected to server.")
    showMenu()
}

func showMenu() {
    var choice string
    for {
        fmt.Println("\n=== Login ===")
        fmt.Println("1. Register")
        fmt.Println("2. Login")
        fmt.Println("0. Exit")
        fmt.Print("Enter your choice: ")
        scanner.Scan()
        choice = scanner.Text()
        handleChoice(choice)
        if choice == "0" {
            break
        }
    }
}

func handleChoice(choice string) {
    switch choice {
    case "1":
        register()
    case "2":
        login()
    case "0":
        fmt.Println("Exiting...")
    default:
        fmt.Println("Invalid choice.")
    }
}

func register() {
    for {
        fmt.Print("Enter username: ")
        scanner.Scan()
        username := scanner.Text()
        fmt.Print("Enter password: ")
        scanner.Scan()
        password := scanner.Text()

        fmt.Fprintln(writer, "REGISTER")
        fmt.Fprintln(writer, username)
        fmt.Fprintln(writer, password)
        writer.Flush()

		response, _ := reader.ReadString('\n')
        if response == "Registration successful!\n" {
            fmt.Println(response)
            createProfile(username)
            break
        } else {
            fmt.Println("Username already exists. Please try again.")
        }
    }
}

func login() {
    fmt.Print("Enter username: ")
    scanner.Scan()
    username := scanner.Text()
    fmt.Print("Enter password: ")
    scanner.Scan()
    password := scanner.Text()

    fmt.Fprintln(writer, "LOGIN")
    fmt.Fprintln(writer, username)
    fmt.Fprintln(writer, password)
    writer.Flush()

    response, _ := reader.ReadString('\n')
    if response == "true\n" {
        loggedInUser = username
        fmt.Println("Login successful!")
        showSocialMediaMenu()
    } else {
        fmt.Println("Login failed. Please try again.")
        showMenu()
    }
}

func showSocialMediaMenu() {
    var choice string
    for {
        fmt.Println("\n=== Social Media Platform Menu ===")
        fmt.Println("1. View News Feed")
        fmt.Println("2. Create Post")
        fmt.Println("3. Comment on Post")
        fmt.Println("4. Search User")
        fmt.Println("5. Upvote Post")
        fmt.Println("6. Downvote Post")
        fmt.Println("0. Logout")
        fmt.Print("Enter your choice: ")
        scanner.Scan()
        choice = scanner.Text()
        handleSocialMediaChoice(choice)
        if choice == "0" {
            break
        }
    }
}

func handleSocialMediaChoice(choice string) {
    switch choice {
    case "1":
        viewNewsFeed()
    case "2":
        createPost()
    case "3":
        commentOnPost()
    case "4":
        searchUser()
    case "5":
        upvotePost()
    case "6":
        downvotePost()
    case "0":
        fmt.Println("Logging out...")
    default:
        fmt.Println("Invalid choice.")
    }
}

func viewNewsFeed() {
    fmt.Fprintln(writer, "VIEW_NEWS_FEED")
    writer.Flush()

    for {
        response, _ := reader.ReadString('\n')
        if response == "\n" {
            break
        }
        fmt.Println(response)
    }
}

func createPost() {
    fmt.Print("Enter post title: ")
    scanner.Scan()
    title := scanner.Text()
    fmt.Print("Enter post content: ")
    scanner.Scan()
    content := scanner.Text()
    fmt.Print("Enter image URL (press Enter to skip): ")
    scanner.Scan()
    imageURL := scanner.Text()

    fmt.Fprintln(writer, "CREATE_POST")
    fmt.Fprintln(writer, title)
    fmt.Fprintln(writer, content)
    fmt.Fprintln(writer, loggedInUser)
    fmt.Fprintln(writer, imageURL)
    writer.Flush()
}

func commentOnPost() {
    fmt.Print("Enter the username whose post you want to comment on: ")
    scanner.Scan()
    postAuthor := scanner.Text()
    fmt.Print("Enter the title of the post you want to comment on: ")
    scanner.Scan()
    postTitle := scanner.Text()
    fmt.Print("Enter your comment: ")
    scanner.Scan()
    comment := scanner.Text()

    fmt.Fprintln(writer, "COMMENT_ON_POST")
    fmt.Fprintln(writer, postAuthor)
    fmt.Fprintln(writer, postTitle)
    fmt.Fprintln(writer, loggedInUser)
    fmt.Fprintln(writer, comment)
    writer.Flush()

    response, _ := reader.ReadString('\n')
    fmt.Println(response)
}

func searchUser() {
    fmt.Print("Enter the username to search for: ")
    scanner.Scan()
    username := scanner.Text()

    fmt.Fprintln(writer, "SEARCH_USER")
    fmt.Fprintln(writer, username)
    writer.Flush()

    response, _ := reader.ReadString('\n')
    fmt.Println(response)
}

func upvotePost() {
    fmt.Print("Enter the title of the post you want to upvote: ")
    scanner.Scan()
    postTitle := scanner.Text()

    fmt.Fprintln(writer, "UPVOTE_POST")
    fmt.Fprintln(writer, postTitle)
    writer.Flush()

    response, _ := reader.ReadString('\n')
    fmt.Println(response)
}

func downvotePost() {
    fmt.Print("Enter the title of the post you want to downvote: ")
    scanner.Scan()
    postTitle := scanner.Text()

    fmt.Fprintln(writer, "DOWNVOTE_POST")
    fmt.Fprintln(writer, postTitle)
    writer.Flush()

    response, _ := reader.ReadString('\n')
    fmt.Println(response)
}

func createProfile(username string) {
    fmt.Print("Enter about: ")
    scanner.Scan()
    about := scanner.Text()
    fmt.Print("Enter experience: ")
    scanner.Scan()
    experience := scanner.Text()
    fmt.Print("Enter education: ")
    scanner.Scan()
    education := scanner.Text()
    fmt.Print("Enter awards: ")
    scanner.Scan()
    awards := scanner.Text()
    fmt.Print("Enter skills: ")
    scanner.Scan()
    skills := scanner.Text()
    fmt.Print("Enter status: ")
    scanner.Scan()
    status := scanner.Text()

    fmt.Fprintln(writer, "CREATE_PROFILE")
    fmt.Fprintln(writer, username)
    fmt.Fprintln(writer, about)
    fmt.Fprintln(writer, experience)
    fmt.Fprintln(writer, education)
    fmt.Fprintln(writer, awards)
    fmt.Fprintln(writer, skills)
    fmt.Fprintln(writer, status)
    writer.Flush()

    response, _ := reader.ReadString('\n')
    fmt.Println(response)
}