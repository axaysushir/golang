package main

import (
	"fmt"
	"sync"
)

// Channels are similar to async await asynchronous operations like js
func main() {
	fmt.Println("Channels in golang")
	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// fmt.Println(<-mych)
	// myCh <- 5
	// use <- to incoming and outgoing channels
	wg.Add(2)

	// Read only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChanelOpen := <-myCh

		fmt.Println(isChanelOpen)
		fmt.Println(val)

		wg.Done()

	}(myCh, wg)

	// Send only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)

		wg.Done()

	}(myCh, wg)
	wg.Wait()
}
