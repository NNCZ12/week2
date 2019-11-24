package main

import (
	"fmt"
	"string"
)

func main() {
	fmt.Println(strings.Count("Hello World", "0"))
	fmt.Println(strings.Count("Hello World", ""))
}