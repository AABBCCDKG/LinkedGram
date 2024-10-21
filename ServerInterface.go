// A program Team Project
//
// Purdue University -- CS18000 -- Spring 2024 -- Team Project
//
// Author: Purdue CS
// Version: Mar 31, 2024

package main

import (
	"net"
	"sync"
)

// ServerInterface defines the methods required for a server in Go
type ServerInterface interface {
	Run()
	GetListener() net.Listener
	GetWaitGroup() *sync.WaitGroup
}

// Server is a struct that implements ServerInterface
type Server struct {
	listener   net.Listener
	waitGroup  *sync.WaitGroup
}

// NewServer creates a new Server instance
func NewServer(listener net.Listener) *Server {
	return &Server{
		listener:  listener,
		waitGroup: &sync.WaitGroup{},
	}
}

// Run starts the server
func (s *Server) Run() {
	// Example: Handle connections here
	defer s.waitGroup.Done()
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			break
		}
		// Handle the connection in a new goroutine
		go s.handleConnection(conn)
	}
}

// GetListener returns the server's Listener (equivalent to ServerSocket)
func (s *Server) GetListener() net.Listener {
	return s.listener
}

// GetWaitGroup returns the server's WaitGroup (equivalent to ExecutorService handling goroutines)
func (s *Server) GetWaitGroup() *sync.WaitGroup {
	return s.waitGroup
}

// handleConnection simulates handling a client connection
func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	// Handle the connection here
}

func main() {
	// Example usage of the ServerInterface implementation
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := NewServer(listener)
	server.waitGroup.Add(1)
	go server.Run()

	// Wait for all goroutines to finish
	server.GetWaitGroup().Wait()
}
