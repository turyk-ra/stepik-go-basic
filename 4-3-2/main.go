package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	result := 1
	for i := a; i <= b; i++ {
		result *= i
	}
	fmt.Print(result)
}
