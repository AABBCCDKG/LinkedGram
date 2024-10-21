package main

import (
    "fmt"
    "log"
    "net"
    "os"
    "runtime"
)

// Server struct
type Server struct{}

// Run method for Server
func (s *Server) Run() {
    // Server logic here
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
    defer listener.Close()
    fmt.Println("Server started.")
}

// ClientGUI struct
type ClientGUI struct{}

// NewClientGUI function to initialize ClientGUI
func NewClientGUI() (*ClientGUI, error) {
    // Client GUI logic here
    fmt.Println("Client GUI started.")
    return &ClientGUI{}, nil
}

func main() {
    // 启动Server
    go func() {
        server := &Server{}
        server.Run()
    }()

    // 启动Client GUI
    go func() {
        _, err := NewClientGUI()
        if err != nil {
            log.Fatalf("Error starting client GUI: %v", err)
        }
    }()

    // Prevent the main function from exiting immediately
    runtime.Goexit()
}