package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.com/robots.txt")
	if err != nil {
		log.Fatalln("Failed to request robots")
	}
	log.Println(resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln("Failed to parse body")
	}
	log.Println(string(body))
}
