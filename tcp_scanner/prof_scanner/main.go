package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int) {
	for port := range ports {
		fmt.Printf("Checking port %d\n", port)
		address := fmt.Sprintf("snanme.nmap.org:%d", port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		results <- port
		conn.Close()
	}
}

func main() {
	ports := make(chan int, 512)
	results := make(chan int)
	var openports []int

	for i := 1; i <= cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}

	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("Port %d is open\n", port)
	}
}
