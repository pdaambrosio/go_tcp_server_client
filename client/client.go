package main

import (
	"fmt"
	"net"
)

// tcpClient is a function that connects to a TCP server and sends a message
func tcpClient(serverAddress string, port int) string {
	// Connect to the server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverAddress, port))

	// Check if there is an error while connecting
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
	}

	defer conn.Close()

	// Send a message to the server
	conn.Write([]byte("Hello, server!"))
	buffer := make([]byte, 1024)

	// Read the response from the server
	conn.Read(buffer)
	return "Response from server: " + string(buffer)
}

func main() {
	// Print the response from the server
	fmt.Println(tcpClient("localhost", 8080))
}
