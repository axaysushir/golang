package main

import "fmt"

func main() {
	fmt.Println("Array in golang")

	var fruits [4]string
	fruits[0] = "go"
	fruits[1] = "Apple"
	fruits[3] = "banana"

	fmt.Println(fruits)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Print(primes)
}

func demo() {
	fmt.Println("")
	var data [4]string
	data[0] = 'oppo'
	data[1] = 'delux'
	data[2] = 'apple'
	data[3] = "data"
	fmt.println(data)
}