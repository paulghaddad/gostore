package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const socketFile = "./unix_sock"

func main() {
	fmt.Println("Creating Unix Socket")
	Run()
}

// Run the gostore server
func Run() {
	os.Remove(socketFile)
	listener, err := net.Listen("unix", socketFile)
	if err != nil {
		log.Fatalf("Problem connecting to Unix socket %s with error: %s", socketFile, err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Problem accepting connection with error: %s", err)
		}
		go HandleConn(conn)
	}
}

// HandleConn handles incoming connections and processes them
func HandleConn(c net.Conn) {
	received := make([]byte, 0)
	for {
		buf := make([]byte, 512)
		count, err := c.Read(buf)
		received = append(received, buf[:count]...)
		if err != nil {
			ProcessMessage(received)
			if err != io.EOF {
				log.Fatalf("Error on read: %s", err)
			}
			break
		}

	}
}

// ProcessMessage handles incoming messages on the server
func ProcessMessage(message []byte) {
	log.Printf("Received a message: %s", message)
}
