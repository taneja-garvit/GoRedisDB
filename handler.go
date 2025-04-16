package main

import (
	"fmt"
	"net"
)

func handleCommand(conn net.Conn, cmd Value) {
	if cmd.typ == "simple" && cmd.str == "PING" {
		fmt.Fprintf(conn, "+PONG\r\n")
	} else {
		fmt.Fprintf(conn, "-ERR unknown command\r\n")
	}
}