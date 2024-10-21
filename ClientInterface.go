// A program Team Project
//
// Purdue University -- CS18000 -- Spring 2024 -- Team Project
//
// Author: Purdue CS
// Version: Mar 31, 2024

package main

import "fmt"

// ClientInterface is a Go interface equivalent to the Java interface.
type ClientInterface interface {
    Run()
}

// Client struct implements the ClientInterface.
type Client struct{}

// Run is the method from ClientInterface.
func (c Client) Run() {
    fmt.Println("Running the client interface.")
}

// main function, similar to public static void main in Java.
func main() {
    var client ClientInterface = Client{}
    client.Run()
}
