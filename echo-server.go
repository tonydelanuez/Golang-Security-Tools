package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// echo is a handler function that echoes received data
func echo(conn net.Conn) {
	defer conn.Close()
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}

func main() {
	// Bind to tcp port on all interfaces
	port := fmt.Sprintf(":%s", os.Args[1])
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("Unable to bind to port: %s", err)
	}

	message := fmt.Sprintf("Listening on 0.0.0.0%s", port)
	log.Println(message)

	for {
		// Wait for connection. Create net.Conn on connection established
		conn, err := listener.Accept()
		log.Println("Receieved connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		conn.Write([]byte("Connected to echo server \n"))

		// Handle the connection using goroutine for concurrency
		go echo(conn)
	}
}
