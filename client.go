package main

import (
	"bufio"
	"fmt"
	"net"
)

func sendCommand(conn net.Conn, cmd string) {
	fmt.Fprintf(conn, cmd)
	response, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Response: ", response)

	// If it's a bulk string, read the actual value too
	if len(response) > 0 && response[0] == '$' && response != "$-1\r\n" {
		val, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Value: ", val)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Sending PING")
	sendCommand(conn, "*1\r\n$4\r\nPING\r\n")

	fmt.Println("\nSending SET mykey hello")
	sendCommand(conn, "*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$5\r\nhello\r\n")

	fmt.Println("\nSending GET mykey")
	sendCommand(conn, "*2\r\n$3\r\nGET\r\n$5\r\nmykey\r\n")
}
