package main

import (
	"bufio"
	"fmt"
	"io"
)


//Handles parsing and serializing RESP (Redis Serialization Protocol) messages, e.g., reading commands like PING or SET key value. Uses bufio for efficient I/O.


type Value struct {
	typ  string // e.g., "simple", "bulk", "array"
	str  string
	bulk string
	arr  []Value
}

type Resp struct {
	reader *bufio.Reader
}

func NewResp(rd io.Reader) *Resp {
	return &Resp{reader: bufio.NewReader(rd)}
}

func (r *Resp) readLine() (string, int, error) {
	line, err := r.reader.ReadString('\n')
	if err != nil {
		return "", 0, err
	}
	line = line[:len(line)-2] // Remove \r\n
	return line, len(line) + 2, nil
}

func (r *Resp) Read() (Value, error) {
	_type, err := r.reader.ReadByte()
	if err != nil {
		return Value{}, err
	}
	switch _type {
	case '+':
		line, _, err := r.readLine()
		if err != nil {
			return Value{}, err
		}
		return Value{typ: "simple", str: line}, nil
	default:
		return Value{}, fmt.Errorf("unsupported RESP type: %v", string(_type))
	}
}