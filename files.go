package main

func main() {
	// Files in go lang

	myfile, err := os.Create("Sample.txt")
	Error(err)
	
}

func Error(err error) {
	if err != nil {
		panic(err)
	}
}