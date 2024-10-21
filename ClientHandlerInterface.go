package main

import "io"

// ClientHandlerInterface defines the methods that any client handler must implement
type ClientHandlerInterface interface {
	// Run starts the client handler logic (similar to Java's Runnable)
	Run()

	// HandleRequest processes the incoming client command and returns the result
	HandleRequest(command string) (string, error)
}
