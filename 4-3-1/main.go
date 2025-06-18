package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)
	var fact = 1
	for i := 1; i <= x; i++ {
		fact *= i
	}
	fmt.Print(fact)
}
