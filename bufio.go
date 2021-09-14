package main

import (
	"bufio"
	"fmt"
	"os"
)

func buffer() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Hello, ")
	fmt.Fprint(w, "world")
	w.Flush()
}
