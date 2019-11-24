package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(strings.Contains("Hello World", "world"))
	fmt.Print(strings.Contains("Hello World", "World"))
}
