package main

import (
	"io"
	"log"
	"net"
)

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "nmap.org:80")
	if err != nil {
		log.Fatalln("Unable to connect destination")
		return
	}
	defer dst.Close()

	go func() {
		if _, err := io.Copy(src, dst); err != nil {
			log.Fatalln(err)
		}
	}()
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln("Unable to start listener")
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Unable to accept connection")
			continue
		}
		go handle(conn)
	}
}
