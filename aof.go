package main

import "os"

//Manages append-only file persistence, logging commands to disk for data durability.

type AOF struct {
	file *os.File
}


func NewAOF(path string) (*AOF, error) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &AOF{file: f}, nil
}

func (a *AOF) Write(cmd Value) error {
	_, err := a.file.WriteString(cmd.str + "\r\n")
	return err
}

func (a *AOF) Close() {
	a.file.Close()
}