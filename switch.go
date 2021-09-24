package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GO runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
		fallthrough
	default:
		fmt.Printf("%s.\n", os)
	}
}
