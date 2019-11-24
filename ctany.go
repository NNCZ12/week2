package main

import (
	"fmt"
	"stings"
)

func main() {
	fmt.Println(strings.ContainsAny("Hello World", "hi"))
	fmt.Println(strings.ContainsAny("Hello World", "Hi"))
}