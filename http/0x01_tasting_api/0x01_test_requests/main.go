package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	content := make([]byte, 4096)

	log.Println("Testing GET method")
	log.Println()
	url_string := "http://google.com/robots.txt"
	r1, err := http.Get(url_string)
	if err != nil {
		log.Fatalln("Failed to test GET: ")
	} else {
		defer r1.Body.Close()
		r1.Body.Read(content)
		log.Println(content)
	}
	log.Println()

	log.Println("Testing HEAD method")
	log.Println()
	r2, err := http.Head(url_string)
	if err != nil {
		log.Fatalln("Failed to test HEAD")
	} else {
		defer r2.Body.Close()
		r2.Body.Read(content)
		log.Println(content)
	}
	log.Println()

	log.Println("Testing POST method")
	log.Println()
	form := url.Values{}
	form.Add("foo", "bar")
	r3, err := http.Post(
		url_string,
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		log.Fatalln("Failed to test POST")
	} else {
		defer r3.Body.Close()
		r3.Body.Read(content)
		log.Println(content)
	}
	log.Println()
}
