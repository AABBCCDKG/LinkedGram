package main

import (
	"fmt"
	"net"
	"sync"
)

// User represents a user entity
type User struct {
	Name string
}

// NewsFeed represents a news feed entity
type NewsFeed struct {
	Posts []string // Simplified for example purposes
}

// ClientHandler handles client connections
type ClientHandler struct {
	Conn   net.Conn
	Users  []User
	NewsFeeds []NewsFeed
}

// Server represents the server entity
type Server struct {
	Port       int
	ServerSock net.Listener
	Users      []User
	NewsFeeds  []NewsFeed
	wg         sync.WaitGroup
}

// NewServer creates a new server
func NewServer(port int) (*Server, error) {
	serverSock, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}
	return &Server{
		Port:       port,
		ServerSock: serverSock,
		Users:      []User{},
		NewsFeeds:  []NewsFeed{{}},
	}, nil
}

// Run starts the server
func (s *Server) Run() {
	defer s.ServerSock.Close()
	for {
		conn, err := s.ServerSock.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			clientHandler := ClientHandler{Conn: conn, Users: s.Users, NewsFeeds: s.NewsFeeds}
			clientHandler.HandleRequest()
		}()
	}
}

// HandleRequest handles the client request
func (ch *ClientHandler) HandleRequest() {
	defer ch.Conn.Close()
	// Add your request handling logic here
	fmt.Fprintf(ch.Conn, "Hello, client connected!\n")
}

func main() {
	server, err := NewServer(8081)
	if err != nil {
		fmt.Printf("Error creating server: %v\n", err)
		return
	}

	fmt.Println("Server started on port 8081")

	server.Run()
	server.wg.Wait()
}