package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	sum := 0
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		if x%2 == 0 && x%3 != 0 {
			sum += x
		}
	}
	fmt.Print(sum)
}
