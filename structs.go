package main

import "fmt"

func main() {
	axay := User{"Axay", "axay@go.dev", true, 5}
	fmt.Println(axay)
}

type User struct {
	Name    string
	Email   string
	Status  bool
	Courses int
}
