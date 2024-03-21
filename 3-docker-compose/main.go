package main

import (
	"fmt"
	"time"
)

func main() {
	// infinite loop
	for {
		fmt.Println("Hello, World! From Docker!")
		// sleep for 1 second
		time.Sleep(1 * time.Second)
	}
}
