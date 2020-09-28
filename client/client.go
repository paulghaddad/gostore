package main

import (
	"log"
	"net"
)

func main() {
	message := []byte("Hello, world!")

	sender := Sender{
		SocketFile: "./unix_sock",
	}

	sender.SendMessage(message)
}

// Sender wraps a logger and socket
type Sender struct {
	SocketFile string
}

// SendMessage connects to the socket and sends a message
func (s *Sender) SendMessage(message []byte) {
	c, err := net.Dial("unix", s.SocketFile)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}

	defer c.Close()
	count, err := c.Write(message)
	if err != nil {
		log.Fatalf("Write Error: %s", err)
	}
	log.Printf("Wrote %d bytes", count)
}
