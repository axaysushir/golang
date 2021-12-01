package main

func main() {
	fmt.Println("Defer in go lang")
	// the order are reversed in defer statements last one comes first

	defer fmt.Println("This comes to last \n")
	defer fmt.Println("This is second last \n")
	defer fmt.Println("This is on top.")

	fmt.Println("Hello")
}