package main

import (
	"fmt"
)

func main() {
	var radius float32
	fmt.Scan(&radius)
	fmt.Print(3.14 * (radius * radius))
}
