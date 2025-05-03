package main

import (
	"fmt"
	"net"
	"os"
)
//Entry point; sets up the TCP server to listen for client connections (port 6379) and routes commands to handlers.

func main() {
	fmt.Println("Starting Redis on port 6379...")
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		os.Exit(1)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Failed to accept connection: %v\n", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	resp := NewResp(conn)

	for {
		cmd, err := resp.Read()
		if err != nil {
			fmt.Fprintf(conn, "-ERR %v\r\n", err)
			return
		}
		handleCommand(conn, cmd)
	}
}
