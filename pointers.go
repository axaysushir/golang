package main

import "fmt"

//  Go has pointers. A pointer holds the memory address of a value.
// The type *T is a pointer to a T value. Its zero value is nil.

func pointers() {
	fmt.Println("Hello World")

	myNumber := 123

	var ptr *int = &myNumber
	println(ptr)  // points out the memory location like: 0xc00008bf30
	println(*ptr) // prints 123

	*ptr = *ptr + 10
	println(*ptr) // add value 10 in pointer and returns 133+
}
