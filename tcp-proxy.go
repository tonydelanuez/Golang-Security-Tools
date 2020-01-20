package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func handle(src net.Conn, host string) {
	conf := &tls.Config{
		//InsecureSkipVerify: true,
	}
	dst, err := tls.Dial("tcp", fmt.Sprintf("%s:443", host), conf)
	if err != nil {
		log.Fatalln("Unable to connect to our unreachable host", err)
	}
	defer dst.Close()

	// Run in goroutine to prevent io.Copy from blocking
	go func() {
		// Copy our source's output to the destination
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()
	// Copy our destination's output back to our source
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	domain := os.Args[1]
	// Listen on port 80
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln("Unable to bind to port", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle(conn, domain)
	}
}
