package main

import "fmt"

func main() {
	fmt.Println("IF ElSE in golang")

	myNum := 12

	if myNum > 10 {
		fmt.Println("Num is greater then 10")
	} else if myNum < 10 {
		fmt.Println("Num is less then 10")
	} else {
		fmt.Println("Num is exactly 12")
	}

	if num := 200; num > 200 {
		fmt.Println("Num is greater")
	} else {
		fmt.Println("Num is exact")
	}
}
