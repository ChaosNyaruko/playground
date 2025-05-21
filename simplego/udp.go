package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Define the target address and port
	targetAddr := "71.18.247.153:3478" // Change this to your target address

	// Resolve the UDP address
	addr, err := net.ResolveUDPAddr("udp", targetAddr)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		os.Exit(1)
	}

	// Create a UDP connection
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error creating connection:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Set a deadline for reading
	// err = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	// if err != nil {
	// 	fmt.Println("Error setting read deadline:", err)
	// 	os.Exit(1)
	// }

	// Send a message
	message := []byte("ello")
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
	fmt.Println("Message sent:", string(message))

	// Receive responses
	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			break
		}
		fmt.Printf("Received %s from %s\n", string(buffer[:n]), addr.String())
	}
}
