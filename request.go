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
	// res, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// body, err := ioutil.ReadAll(res.Body)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// data := string(body)
	// log.Printf(data)

	postBody, _ := json.Marshal(map[string]string{
		"name":  "James",
		"email": "james@google.com",
	})

	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("https://postman-echo.com/post", "application/json", responseBody)

	if err != nil {
		log.Fatal("An error occured %v ", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	sb := string(body)
	log.Printf(sb)
	// 2021/10/05 13:11:33 {"args":{},"data":{"email":"james@google.com","name":"James"},"files":{},"form":{},"headers":{"x-forwarded-proto":"https","x-forwarded-port":"443","host":"postman-echo.com","x-amzn-trace-id":"Root=1-615c01ad-6734c6eb54e268eb74643092","content-length":"43","content-type":"application/json","accept-encoding":"gzip","user-agent":"Go-http-client/2.0"},"json":{"email":"james@google.com","name":"James"},"url":"https://postman-echo.com/post"}
}
