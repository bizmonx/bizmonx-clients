package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to the TCP server on localhost:8080
	conn, err := net.Dial("tcp", "localhost:1984")
	if err != nil {
		// Handle error
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	// Close the connection when the function exits
	defer conn.Close()

	// Message to send
	message := "status newhost.go green message from go client\n"

	// Send the message
	_, err = conn.Write([]byte(message))
	if err != nil {
		// Handle error
		fmt.Println("Error sending message:", err.Error())
		return
	}

	fmt.Println("Message sent successfully")
}
