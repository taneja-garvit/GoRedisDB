package main

import (
	"fmt"
	"net"
	"strings"
)

var store = make(map[string]string)

func handleCommand(conn net.Conn, cmd Value) {
	if cmd.typ == "array" && len(cmd.arr) > 0 {
		command := strings.ToUpper(cmd.arr[0].str)

		switch command {
		case "PING":
			fmt.Fprintf(conn, "+PONG\r\n")
		case "SET":
			if len(cmd.arr) != 3 {
				fmt.Fprintf(conn, "-ERR wrong number of arguments for 'SET'\r\n")
				return
			}
			key := cmd.arr[1].str
			value := cmd.arr[2].str
			store[key] = value
			fmt.Fprintf(conn, "+OK\r\n")
		case "GET":
			if len(cmd.arr) != 2 {
				fmt.Fprintf(conn, "-ERR wrong number of arguments for 'GET'\r\n")
				return
			}
			key := cmd.arr[1].str
			if val, ok := store[key]; ok {
				fmt.Fprintf(conn, "$%d\r\n%s\r\n", len(val), val)
			} else {
				fmt.Fprintf(conn, "$-1\r\n") // Key not found
			}
		default:
			fmt.Fprintf(conn, "-ERR unknown command '%s'\r\n", command)
		}
	} else {
		fmt.Fprintf(conn, "-ERR invalid command format\r\n")
	}
}
