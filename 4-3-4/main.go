package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	result := 1
	for i := 1; i <= x; i++ {
		if i%2 == 0 {
			result *= i
		}
	}
	fmt.Print(result)
}
