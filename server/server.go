package main

import (
	"fmt"
	"net"
)

func tcpServer(conn net.Conn) {
	defer conn.Close() // This will be executed at the end of the function (close the connection)
	buffer := make([]byte, 1024) // Create a buffer of 1024 bytes to read the incoming data
	_, err := conn.Read(buffer) // Read the incoming data and put it in the buffer

	// Check if there is an error while reading the data from the connection (client)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	// Print the incoming data
	fmt.Println("Received data:", string(buffer))

	// Send a response to the client
	conn.Write([]byte("Message received."))
}

func main() {
	// Listen on port 8080
	ln, err := net.Listen("tcp", ":8080")

	// Check if there is an error while listening
	if err != nil {
		fmt.Println("Error listening:", err.Error())
	}

	// Close the listener at the end of the function
	defer ln.Close()

	// Print a message to the console
	fmt.Println("Listening on port 8080")

	// Infinite loop to accept new connections
	for {
		// Accept a new connection
		conn, err := ln.Accept()

		// Check if there is an error while accepting a new connection
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
		}

		// Handle the new connection
		go tcpServer(conn)
	}
}
