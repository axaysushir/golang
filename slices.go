// slices in golang
package main

import "fmt"

func main() {
	fmt.Println("Slices in golang")

	slices := []string{"hello", "World", "Go", "lang"}
	fmt.Println(slices[0:4]) // prints index 0 to 3 excludes 4 [hello World Go lang]

	names := [4]string{"hello", "World", "Go", "lang"}

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b) // [hello World] [World Go]

	b[0] = "ABC"
	fmt.Println(a, b)  // [hello ABC] [ABC Go]
	fmt.Println(names) // [hello ABC Go lang]

	//  The length of a slice is the number of elements it contains.

	// The capacity of a slice is the number of elements in the underlying array,
	// counting from the first element in the slice.

	// The length and capacity of a slice s can be obtained using the expressions
	// len(s) and cap(s).
	// The zero value of a slice is nil
	fmt.Println(len(names), cap(names))

	// slices with make

	mySlice := make([]int, 5)

	myOtherSlice := make([]int, 0, 5)  // add capacity by passing third argument
	fmt.Println(myOtherSlice, mySlice) // [] [0 0 0 0 0]

	// Append method

	var s []int
	s = append(s, 0)
	s = append(s, 1, 2, 4)
	printSlice(s) // len=2 cap=2 [0 1 2 4]

	// Ranges in slice
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	// 2**0 = 1
	// 2**1 = 2
	// 2**2 = 4
	// 2**3 = 8
	// 2**4 = 16
	// 2**5 = 32
	// 2**6 = 64
	// 2**7 = 128

}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
