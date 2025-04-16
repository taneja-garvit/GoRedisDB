package main

import (
	"fmt"
	"net"
)
// Implements command logic (e.g., PING returns PONG, SET stores a key-value pair, GET retrieves it)3q2

func handleCommand(conn net.Conn, cmd Value) {
	if cmd.typ == "simple" && cmd.str == "PING" {
		fmt.Fprintf(conn, "+PONG\r\n")
	} else {
		fmt.Fprintf(conn, "-ERR unknown command\r\n")
	}
}