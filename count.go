package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Count("Hello World", "0"))
	fmt.Println(strings.Count("Hello World", ""))
}
