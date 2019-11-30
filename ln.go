package main

import "fmt"

func main() {
	n, e := fmt.Println("Hello", "World", 123, 456)
	fmt.Println("number of bytes written :", n, "\n")
	fmt.Println("write error encountered :", e, "\n")
}
