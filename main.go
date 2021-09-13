package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello")
	today := time.Now()
	fmt.Println(today.Format("01-02-2006 15:05 Monday"))
}
