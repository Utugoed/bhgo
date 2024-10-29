package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Start testing...")

	req, err := http.NewRequest("DELETE", "http://google.com/robots.txt", nil)
	if err != nil {
		log.Fatalln("Failed to create HTTP Request")
		return
	}
	var client http.Client
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln("Failed to delete robots")
		return
	}
	defer res.Body.Close()
	content := make([]byte, 4096)
	res.Body.Read(content)
	log.Println(content)
}
