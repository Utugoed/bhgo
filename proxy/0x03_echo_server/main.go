package main

import (
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 512)
	for {
		# Здесь переменную input_len видно во всей функции
		input_len, err := conn.Read(buffer)
		if err == io.EOF {
			log.Println("Client has disconnected")
			break
		}
		# Даже после иф
		if err != nil {
			log.Println("Unknown error")
			break
		}
		log.Printf("Received %d bytes", input_len)

		if a, err := conn.Write(buffer[0:input_len]); err != nil {
			# А а видно только внутри иф
			log.Println("Error occured during the echoing")
		}
	}
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
