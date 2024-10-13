package main

import (
	"fmt"
	"io"
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

func Copy(src io.Reader, dst io.Writer) (int, error) {
	input := make([]byte, 4096)

	input_len, err := src.Read(input)
	if err != nil {
		log.Fatalln("Error occured during the reading input!")
		return 0, err
	}
	fmt.Printf("Custom reader have read %d bytes\n", input_len)

	written_len, err := dst.Write(input)
	if err != nil {
		log.Fatalln("Error occured during the writing input")
		return 0, err
	}
	fmt.Printf("Custom writer have written %d bytes\n", written_len)
	return input_len, nil
}

func main() {
	var (
		reader CustomReader
		writer CustomWriter
	)

	_, err := Copy(&reader, &writer)
	if err != nil {
		log.Fatalf("Error occured during the copying")
	}
}
