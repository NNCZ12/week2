package main

import (
	"fmt"
	"time"
)

func main() {
	startA := time.Now()
	var a string
	for i := 0; i < 10000; i++ {
		a += "a"
	}
	fmt.Println("a", time.Since(startA))

	startB := time.Now()
	var a string.Builder
	for i := 0; i < 10000; i++ {
		b.WriteString("b")
	}
	fmt.Println("b", time.Since(startB))
}
