# Golang Security Tools

Tools used for penetration testing adapted from the book [Black Hat Go](https://nostarch.com/blackhatgo). 


[scanner.go](./scanner.go): Simple TCP port scanner. Utilizes goroutines to perform scans concurrently. 

[echo-server.go](./echo-server.go): TCP echo server that listens on a given port and echoes back every message sent.

[tcp-proxy.go](./tcp-proxy.go): TCP Proxy to run on a local/owned webserver and fetch remote resources.