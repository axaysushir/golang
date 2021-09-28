package main

import "fmt"

// functions in golang

func main() {
	fmt.Println("Functions")
	data := adder(12, 4)
	fmt.Println(data)
}

func adder(a int, b int) int {
	return a + b
}
