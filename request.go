package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// web request in golang
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	data := string(body)
	log.Printf(data)

	postBody, _ := json.Marshal(map[string]string{
		"name":  "James",
		"email": "james@google.com",
	})

	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("https://postman-echo.com/post", "application/json", responseBody)

	if err != nil {
		log.Fatal("An error occured %v ", err, resp)
	}
}
