package main

import (
	"fmt"
	"log"
	"net"
)

func test_ip(i1, i2, i3, i4 int) (bool, error) {
	fmt.Printf("Checking %d.%d.%d.%d", i1, i2, i3, i4)
	address := fmt.Sprintf("%d.%d.%d.%d:80", i1, i2, i3, i4)
	_, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalf("Failed to connect %s", address)
		return false, nil
	}
	return true, nil
}

func main() {
	for i1 := 1; i1 <= 255; i1++ {
		for i2 := 1; i2 <= 255; i2++ {
			for i3 := 1; i3 <= 255; i3++ {
				for i4 := 1; i4 <= 255; i4++ {

				}
			}
		}
	}
}
