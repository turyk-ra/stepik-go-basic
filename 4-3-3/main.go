package main

import (
	"fmt"
)

func main() {
	var a, b int
	result := 1
	fmt.Scan(&a, &b)
	for i := a; i <= b; i++ {
		if i%10 == 7 || i%10 == -7 {
			result *= i
		}
	}
	fmt.Print(result)
}
