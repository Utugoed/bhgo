package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	content := make([]byte, 4096)
	form := url.Values{}
	form.Add("foo", "bar")
	req, err := http.NewRequest(
		"PUT", "http://google.com/robots.txt", strings.NewReader(form.Encode()),
	)
	if err != nil {
		log.Fatalln("Failed to create an HTTP Request")
		return
	}
	var client http.Client
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln("Failed to send a request to host")
		return
	}
	defer res.Body.Close()
	res.Body.Read(content)
	log.Println(string(content[:]))
}
