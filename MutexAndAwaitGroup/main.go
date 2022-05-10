package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition in go with mutex and await group")

	wg := &sync.WaitGroup{}
	mutex := &sync.RWMutex{} // read write mutex

	var score = []int{0}
	wg.Add(3) // add  number of func to execute
	// use lock and unlock to prevent mishandling
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One Read")
		mutex.Lock()
		score = append(score, 1)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Tow R")
		mutex.Lock()
		score = append(score, 2)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		mutex.Lock()
		score = append(score, 3)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Four R")
		mutex.RLock()
		fmt.Println(score)
		mutex.RUnlock()
		wg.Done()
	}(wg, mutex)

	wg.Wait()
	fmt.Println(score)
}

// Output:
// Race condition in go
// One R
// Tow R
// Three R
// Four R
// [0 1 2 3]
// [0 1 2 3]
