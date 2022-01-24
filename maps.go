package main

import "fmt"

func main() {
	lists := make(map[int]string)
	lists[0] = "JavaScript"
	lists[1] = "Python"
	lists[2] = "Java"
	lists[3] = "Ruby"

	fmt.Println(lists) // map[0:JavaScript 1:Python 2:Java 3:Ruby]

	// delete
	delete(lists, 2)   // pass list and key as params
	fmt.Println(lists) // map[0:JavaScript 1:Python 3:Ruby]

	for key, value := range lists {
		fmt.Printf("Key: %d, Value: %v ", key, value) // Key: 3, Value: Ruby Key: 0, Value: JavaScript Key: 1, Value: Python
	}

	data := make(map[int]string)
	data[0] = "js"
	data[1] = "py"
	data[2] = "Slang"
	delete(data, 2)
	fmt.Println(data)

	for key, value := range data {
		fmt.Printf("key: %d, value: %v", key, value)
	}
}
