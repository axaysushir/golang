package main

import "fmt"

func main() {
	// lang := []string{"JavaScript", "Python", "Go", "Dart"}

	// for _, i := range lang {
	// 	fmt.Println(i)
	// }

	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum) // 45

	// for sum < 1000 {
	// 	sum += sum
	// }
	// fmt.Println(sum) // 1440
	sum := 1234
	for {
		sum *= 1234
		// fmt.Println(sum) // stops at 3006029
	}
	fmt.Println(sum)
}
