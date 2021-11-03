package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	// go greeter("Hello")
	// greeter("World")
	weblist := []string{
		"https://google.com",
		"https://github.com",
		"https://fb.com",
	}

	for _, web := range weblist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(web string) {
	res, err := http.Get(web)

	if err != nil {
		fmt.Println("err in endpoint")
	} else {
		mutex.Lock()
		signals = append(signals, web)
		mutex.Unlock()
		fmt.Printf("%d status code %s\n", res.StatusCode, web)
	}
}
