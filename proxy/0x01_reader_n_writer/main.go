package main

import (
	"fmt"
	"log"
	"os"
)

type CustomReader struct{}

func (reader *CustomReader) Read(buffer []byte) (int, error) {
	fmt.Print("Input > ")
	return os.Stdin.Read(buffer)
}

type CustomWriter struct{}

func (writer *CustomWriter) Write(buffer []byte) (int, error) {
	fmt.Print("Output > ")
	return os.Stdout.Write(buffer)
}

func main() {
	var (
		reader CustomReader
		writer CustomWriter
	)

	input := make([]byte, 4096)
	input_len, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Error occured during the reading input!")
	}
	fmt.Printf("Custom reader have read %d bytes\n", input_len)

	written_len, err := writer.Write(input)
	if err != nil {
		log.Fatalln("Error occured during the writing input")
	}
	fmt.Printf("Custom writer have written %d bytes\n", written_len)
}
