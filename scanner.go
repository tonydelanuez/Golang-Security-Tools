package main

import (
	"fmt"
	"net"
	"os"
	"sort"
	"time"
)

func connect(host string, ports chan int, results chan int) {
	for port := range ports {
		dest := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", dest, 5*time.Second)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		// Pass open ports back to the channel
		results <- port
	}
}

func main() {
	ip := os.Args[1]

	message := fmt.Sprintf("Starting scan on host %s", ip)
	fmt.Println(message)
	numWorkers := 100
	ports := make(chan int, numWorkers)
	results := make(chan int)
	portLimit := 65355
	var openports []int

	for i := 0; i <= cap(ports); i++ {
		go connect(ip, ports, results)
	}

	/* Start a separate goroutine to pass ports
	to the ports channel. Result-gathering loop needs
	to be able to start before work can continue */
	go func() {
		for i := 1; i <= portLimit; i++ {
			ports <- i
		}
	}()

	/* Receive on the results channel
	for every possible port */
	for i := 0; i < portLimit; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openports)

	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
