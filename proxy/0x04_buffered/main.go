package main

import (
	"bufio"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(s); err != nil {
		log.Fatalln("Unable to write data")
	}
	writer.Flush()
}

func main() {
	listener, err := net.Listen("tcp", ":28800")
	if err != nil {
		log.Println("An error occured during the listen starting")
	}
	log.Printf("Listening on port 28800")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error occured during the connection accepting")
		}
		go echo(conn)
	}
}
